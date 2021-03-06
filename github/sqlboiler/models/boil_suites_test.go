// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Emails", testEmails)
	t.Run("EmailSends", testEmailSends)
	t.Run("EmailTimers", testEmailTimers)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Emails", testEmailsDelete)
	t.Run("EmailSends", testEmailSendsDelete)
	t.Run("EmailTimers", testEmailTimersDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Emails", testEmailsQueryDeleteAll)
	t.Run("EmailSends", testEmailSendsQueryDeleteAll)
	t.Run("EmailTimers", testEmailTimersQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Emails", testEmailsSliceDeleteAll)
	t.Run("EmailSends", testEmailSendsSliceDeleteAll)
	t.Run("EmailTimers", testEmailTimersSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Emails", testEmailsExists)
	t.Run("EmailSends", testEmailSendsExists)
	t.Run("EmailTimers", testEmailTimersExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Emails", testEmailsFind)
	t.Run("EmailSends", testEmailSendsFind)
	t.Run("EmailTimers", testEmailTimersFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Emails", testEmailsBind)
	t.Run("EmailSends", testEmailSendsBind)
	t.Run("EmailTimers", testEmailTimersBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Emails", testEmailsOne)
	t.Run("EmailSends", testEmailSendsOne)
	t.Run("EmailTimers", testEmailTimersOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Emails", testEmailsAll)
	t.Run("EmailSends", testEmailSendsAll)
	t.Run("EmailTimers", testEmailTimersAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Emails", testEmailsCount)
	t.Run("EmailSends", testEmailSendsCount)
	t.Run("EmailTimers", testEmailTimersCount)
	t.Run("Users", testUsersCount)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Emails", testEmailsInsert)
	t.Run("Emails", testEmailsInsertWhitelist)
	t.Run("EmailSends", testEmailSendsInsert)
	t.Run("EmailSends", testEmailSendsInsertWhitelist)
	t.Run("EmailTimers", testEmailTimersInsert)
	t.Run("EmailTimers", testEmailTimersInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Emails", testEmailsReload)
	t.Run("EmailSends", testEmailSendsReload)
	t.Run("EmailTimers", testEmailTimersReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Emails", testEmailsReloadAll)
	t.Run("EmailSends", testEmailSendsReloadAll)
	t.Run("EmailTimers", testEmailTimersReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Emails", testEmailsSelect)
	t.Run("EmailSends", testEmailSendsSelect)
	t.Run("EmailTimers", testEmailTimersSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Emails", testEmailsUpdate)
	t.Run("EmailSends", testEmailSendsUpdate)
	t.Run("EmailTimers", testEmailTimersUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Emails", testEmailsSliceUpdateAll)
	t.Run("EmailSends", testEmailSendsSliceUpdateAll)
	t.Run("EmailTimers", testEmailTimersSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
