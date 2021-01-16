package main

import (
	"bytes"
	"encoding/binary"
	"github.com/imroc/biu"
	"strconv"
)

type dnsHeader struct {
	id      []byte //��ʶ�ֶΣ��ͻ��˻�������������ص�DNSӦ���ģ���ȡIDֵ�����������õ�IDֵ���Ƚϣ������ͬ������Ϊ��ͬһ��DNS�Ự��
	qr      int64  //0��ʾ��ѯ���ģ�1��ʾ��Ӧ����;
	opcode  int64  //0����׼��ѯ��������ֵΪ1�������ѯ����2��������״̬����,[3,15]����ֵ;
	aa      int64  //��ʾ��Ȩ�ش�authoritative answer��-- �������λ��Ӧ���ʱ��������壬ָ������Ӧ��ķ������ǲ�ѯ��������Ȩ����������;
	tc      int64  //��ʾ�ɽضϵģ�truncated��--����ָ�����ı�����ĳ��Ȼ�Ҫ�������±��ض�;
	rd      int64  //��ʾ�����ݹ�(Recursion Desired) -- �������λ���������ã�Ӧ���ʱ��ʹ�õ���ͬ��ֵ���ء����������RD���ͽ����������������еݹ�������ݹ��ѯ��֧���ǿ�ѡ��;
	ra      int64  //��ʾ֧�ֵݹ�(Recursion Available) -- �������λ��Ӧ�������û�ȡ������������������Ƿ�֧�ֵݹ��ѯ;
	z       int64  //����ֵ����δʹ��;
	rcode   int64  //0 : û�д���;1 : ���ĸ�ʽ����(Format error);2 : ������ʧ��(Server failure);3 : ����������(Name Error);4 : ������������֧�ֲ�ѯ����(Not Implemented);5 : �ܾ�(Refused)
	qdcount int64  //����������е������¼��
	ancount int64  //���Ļش���еĻش��¼��
	nscount int64  //������Ȩ���е���Ȩ��¼��
	arcount int64  //���ĸ��Ӷ��еĸ��Ӽ�¼��
}

type dnsQuestion struct {
	qname  []byte //����
	qtype  []byte //��ѯ��Э������
	qclass []byte //��ѯ����,���磬IN����Internet
}

type dnsAnswer struct {
	aname  []byte //����
	atype  []byte //��ѯ��Э������
	aclass []byte //��ѯ����,���磬IN����Internet
	ttl    int64  //time to live���ʱ��
	rdlen  int64  //���ݳ���
	rdata  []byte //���ݼ�¼
}

type dnsRespone struct {
	header   dnsHeader
	question dnsQuestion
	answer   dnsAnswer
}

var qtypeDic = map[int64]string{
	1:   "A",     //IPv4��ַ
	2:   "NS",    //���ַ�����
	5:   "CNAME", //�淶���ƶ�����������ʽ���ֵı���
	6:   "SOA",   //��ʼ��Ȩ���һ�����Ŀ�ʼ
	11:  "WKS",   //��֪�����������ṩ���������
	12:  "PTR",   //ָ���IP��ַת��Ϊ����
	13:  "HINFO", //������Ϣ��������ʹ�õ�Ӳ���Ͳ���ϵͳ�ı���
	15:  "MX",    //�ʼ��������ʼ��ı�·���͵��ʼ�������
	28:  "AAAA",  //IPv6��ַ
	252: "AXFR",  //����������������
	255: "ANY",   //�����м�¼������
}

func getHeader(data []byte) dnsHeader {
	flags := data[2:4]

	flagsStr := biu.BytesToBinaryString(flags)
	//log.Printf("flags:%x  flagsBit: %q ",flags,flagsStr)

	header := &dnsHeader{
		id:      data[0:2],
		qr:      StringToInt64(flagsStr[1:2]),
		opcode:  StringToInt64(flagsStr[2:6]),
		aa:      StringToInt64(flagsStr[6:7]),
		tc:      StringToInt64(flagsStr[7:8]),
		rd:      StringToInt64(flagsStr[8:9]),
		ra:      StringToInt64(flagsStr[10:11]),
		z:       StringToInt64(flagsStr[11:14]),
		rcode:   StringToInt64(flagsStr[14:18]),
		qdcount: BytesToInt64(data[4:6]),
		ancount: BytesToInt64(data[6:8]),
		nscount: BytesToInt64(data[6:8]),
		arcount: BytesToInt64(data[10:12]),
	}

	return *header
}

func getQuestion(data []byte) dnsQuestion {

	qnameBytes := data[12 : len(data)-4]

	question := &dnsQuestion{
		qname:  qnameBytes,
		qtype:  data[len(data)-4 : len(data)-2], //BytesToInt64(data[len(data)-4:len(data)-2]),
		qclass: data[len(data)-2 : len(data)],   //BytesToInt64(data[len(data)-2:len(data)]),
	}

	return *question
}

//int64ת byte ����
func Int64ToBytes(i int64) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, i)
	return b_buf.Bytes()[len(b_buf.Bytes())-2:]
}

//byte ����ת int64
func BytesToInt64(buf []byte) int64 {
	bufStr := biu.BytesToBinaryString(buf)
	bufint64, _ := strconv.ParseInt(bufStr[len(bufStr)-5:len(bufStr)-1], 10, 64)
	return bufint64
}

//�ַ���ת int64
func StringToInt64(str string) int64 {
	intStr, _ := strconv.ParseInt(str, 10, 64)
	return intStr
}

func main() {

}
