module github.com/tektoncd/results

go 1.13

require (
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.9.0
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/jackc/pgconn v1.11.0
	github.com/jonboulle/clockwork v0.2.3
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/prometheus/client_golang v1.11.1
	github.com/spf13/viper v1.10.0
	github.com/tektoncd/pipeline v0.40.2
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/zap v1.23.0
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	google.golang.org/api v0.74.0
	google.golang.org/genproto v0.0.0-20220324131243-acbaeb5b85eb
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/postgres v1.3.4
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.4
	k8s.io/api v0.23.9
	k8s.io/apimachinery v0.23.9
	k8s.io/client-go v0.23.9
	knative.dev/pkg v0.0.0-20220818004048-4a03844c0b15
	sigs.k8s.io/yaml v1.3.0
)
