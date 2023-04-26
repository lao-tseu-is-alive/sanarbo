package trees

import (
	"fmt"
	"runtime"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/config"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/database"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/golog"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/tools"
	"github.com/lao-tseu-is-alive/sanarbo/pkg/version"
)

const (
	defaultDBPort              = 5432
	defaultDBIp                = "127.0.0.1"
	defaultDBSslMode           = "prefer"
)

type WorkingEnv struct {
	l golog.MyLogger
	dbConn database.DB
	storage Storage
}

func TestSearchTreesByName(t *testing.T) {
	type args struct {
		t string
	}

	var w = WorkingEnv{}
	err := w.GetLogger(golog.DebugLevel, fmt.Sprintf("%s ", version.APP))
	if err != nil {
		t.Fatalf("got error when getting logger, err: %v", err)
	}
	err = w.GetDb()
	if err != nil {
		t.Fatalf("got error when getting db, err: %v", err)
	}
	err = w.GetStorage()
	if err != nil {
		t.Fatalf("got error when getting storage, err: %v", err)
	}
	defer w.dbConn.Close()

	tests := []struct {
		name		string
		args		args
		wantRes		[]*TreeList
		wantErr		error
	}{
		{
			name: "it should return an object with name attribute matching pattern containing *",
			args: 	args{t: "*Tre*"},
			wantRes: []*TreeList{{Name: "MyNewTree"}},
			wantErr: nil,
		},
		{
			name: "it should return an object with name attribute matching pattern containing %",
			args: 	args{t: "%%Tre%%"},
			wantRes: []*TreeList{{Name: "MyNewTree"}},
			wantErr: nil,
		},
		{
			name: "should return an error if no tree found",
			args: args{t: "666"},
			wantRes: nil,
			wantErr: ErrNoRecordFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotErr := w.storage.SearchTreesByName(tt.args.t)

			if tt.wantErr != nil {
				assertError(t, gotErr, tt.wantErr)
			} else {
				assertNoError(t, gotErr)
				assertName(t, *gotRes[0], tt.wantRes[0].Name)
			}
		})
	}
}

func (w *WorkingEnv) GetLogger(level golog.Level, prefix string) error {
	log, err := golog.NewLogger("zap", level, prefix)
	if err != nil {
		return fmt.Errorf("got no logger error: %v", err)
	} else {
		w.l = log
	}
	return nil
}

func (w *WorkingEnv) GetDb() error {
	dbDsn, err := config.GetPgDbDsnUrlFromEnv(defaultDBIp, defaultDBPort,
		tools.ToSnakeCase(version.APP), version.AppSnake, defaultDBSslMode)
	if err != nil {
		return fmt.Errorf("error doing config.GetPgDbDsnUrlFromEnv error: %v\n", err)
	}
	var dbConn database.DB
	dbConn, err = database.GetInstance("pgx", dbDsn, runtime.NumCPU(), w.l)
	if err != nil {
		return fmt.Errorf("error doing database.GetInstance(\"pgx\", dbDsn)  : %v\n", err)
	} else {
		w.dbConn = dbConn
	}
	return nil
}

func (w *WorkingEnv) GetStorage() error {
	treesStorage, err := GetStorageInstance("pgx", w.dbConn, w.l)
	if err != nil {
		return fmt.Errorf("error doing GetStorageInstance(\"pgx\", %#v, %#v) error: %v\n", w.dbConn, w.l, err)
	} else {
		w.storage = treesStorage
	}
	return nil
}

func assertName(t testing.TB, list TreeList, want string) {
	t.Helper()
	got := list.Name

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}