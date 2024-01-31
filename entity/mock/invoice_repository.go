package mock

import "database/sql"

type MockDB struct {
	MockPrepare     func(*sql.Stmt, error)
	MockExec        func(sql.Result, error)
	MockQueryRow    func(*sql.Row)
	MockScan        func(args ...interface{})
	PrepareCounter  int
	ExecCounter     int
	QueryRowCounter int
	ScanCounter     int
}

func (db *MockDB) Prepare(query string) (*sql.Stmt, error) {
	db.PrepareCounter++
	stmt := &sql.Stmt{}
	if db.MockPrepare != nil {
		db.MockPrepare(stmt, nil)
	}
	return stmt, nil
}

func (db *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	db.ExecCounter++
	result := &MockResult{}
	if db.MockExec != nil {
		db.MockExec(result, nil)
	}
	return result, nil
}

func (db *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	db.QueryRowCounter++
	row := &sql.Row{}
	if db.MockQueryRow != nil {
		db.MockQueryRow(row)
	}
	return row
}

func (db *MockDB) Scan(dest ...interface{}) {
	db.ScanCounter++
	if db.MockScan != nil {
		db.MockScan(dest...)
	}
}

type MockResult struct {
	RowsAffecteds int64
}

func (r *MockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *MockResult) RowsAffected() (int64, error) {
	return r.RowsAffecteds, nil
}

type MockTx struct {
	db *MockDB
}

func (tx *MockTx) Commit() error {
	return nil
}

func (tx *MockTx) Rollback() error {
	return nil
}

func (tx *MockTx) Prepare(query string) (*sql.Stmt, error) {
	return tx.db.Prepare(query)
}

func (tx *MockTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.db.Exec(query, args...)
}
