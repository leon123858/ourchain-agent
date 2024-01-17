package sqlite

import (
	"database/sql"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func TestBeginTx(t *testing.T) {
	client := setUp()
	tx, err := BeginTx(&client)
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("tx is nil")
	}
	err = RollbackTx(tx)
	if err != nil {
		return
	}
	err = CommitTx(tx)
	if err != nil {
		return
	}
	tearDown(client)
}

func TestBlockCreate(t *testing.T) {
	client := setUp()

	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	item := Block{
		Height: 1,
		Hash:   "hash",
	}
	_, err = BlockCreateExec(stmt, item)
	assert.Equal(t, err, nil)
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 0)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	blocks, err = client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)

	tearDown(client)
}

func TestBlockCreateList(t *testing.T) {
	client := setUp()

	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	for i := 1; i < 10; i++ {
		item := Block{
			Height: uint64(i),
			Hash:   "hash",
		}
		_, err = BlockCreateExec(stmt, item)
		assert.Equal(t, err, nil)
	}
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 0)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	blocks, err = client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)
	assert.Equal(t, blocks[0].Height, uint64(9))

	tearDown(client)
}

func TestBlockDelete(t *testing.T) {
	client := setUp()
	// insert
	tx, err := BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err := BlockCreatePrepare(tx)
	assert.Equal(t, err, nil)
	for i := 1; i < 10; i++ {
		item := Block{
			Height: uint64(i),
			Hash:   "hash",
		}
		_, err = BlockCreateExec(stmt, item)
		assert.Equal(t, err, nil)
	}
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// delete
	tx, err = BeginTx(&client)
	assert.Equal(t, err, nil)
	stmt, err = BlockDeletePrepare(tx)
	assert.Equal(t, err, nil)
	item := Block{
		Height: 9,
	}
	_, err = BlockDeleteExec(stmt, item)
	assert.Equal(t, err, nil)
	err = CommitTx(tx)
	assert.Equal(t, err, nil)
	// check
	blocks, err := client.GetFirstBlockInfo()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(blocks), 1)
	assert.Equal(t, blocks[0].Height, uint64(8))
	tearDown(client)
}

func TestExecPrepare(t *testing.T) {
	type args struct {
		stmt *sql.Stmt
		args []interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult sql.Result
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ExecPrepare(tt.args.stmt, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecPrepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ExecPrepare() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPrepareTx(t *testing.T) {
	type args struct {
		tx  *sql.Tx
		sql string
	}
	tests := []struct {
		name     string
		args     args
		wantStmt *sql.Stmt
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStmt, err := PrepareTx(tt.args.tx, tt.args.sql)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStmt, tt.wantStmt) {
				t.Errorf("PrepareTx() gotStmt = %v, want %v", gotStmt, tt.wantStmt)
			}
		})
	}
}

func TestRollbackTx(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RollbackTx(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("RollbackTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUtxoCreateExec(t *testing.T) {
	type args struct {
		stmt *sql.Stmt
		item Utxo
	}
	tests := []struct {
		name       string
		args       args
		wantResult sql.Result
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := UtxoCreateExec(tt.args.stmt, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoCreateExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UtxoCreateExec() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestUtxoCreatePrepare(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name     string
		args     args
		wantStmt *sql.Stmt
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStmt, err := UtxoCreatePrepare(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoCreatePrepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStmt, tt.wantStmt) {
				t.Errorf("UtxoCreatePrepare() gotStmt = %v, want %v", gotStmt, tt.wantStmt)
			}
		})
	}
}

func TestUtxoDeleteExec(t *testing.T) {
	type args struct {
		stmt *sql.Stmt
		item Utxo
	}
	tests := []struct {
		name       string
		args       args
		wantResult sql.Result
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := UtxoDeleteExec(tt.args.stmt, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoDeleteExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UtxoDeleteExec() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestUtxoDeletePrepare(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name     string
		args     args
		wantStmt *sql.Stmt
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStmt, err := UtxoDeletePrepare(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoDeletePrepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStmt, tt.wantStmt) {
				t.Errorf("UtxoDeletePrepare() gotStmt = %v, want %v", gotStmt, tt.wantStmt)
			}
		})
	}
}

func TestUtxoUpdateExec(t *testing.T) {
	type args struct {
		stmt *sql.Stmt
		item Utxo
	}
	tests := []struct {
		name       string
		args       args
		wantResult sql.Result
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := UtxoUpdateExec(tt.args.stmt, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoUpdateExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UtxoUpdateExec() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestUtxoUpdatePrepare(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name     string
		args     args
		wantStmt *sql.Stmt
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStmt, err := UtxoUpdatePrepare(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UtxoUpdatePrepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStmt, tt.wantStmt) {
				t.Errorf("UtxoUpdatePrepare() gotStmt = %v, want %v", gotStmt, tt.wantStmt)
			}
		})
	}
}
