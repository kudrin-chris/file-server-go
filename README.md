# file-server-go
Сервер, который берет конфиги для соединения с бд, (схема в "./interal/app/model/diagramDB/..." 
Парсер, который получает список .tsv файлов из директории "./inputfile", парсит данные, кладет их в соответствующие структуры(для записи в бд), 
файл с данными по каждому устройству кладется в директорию "./outputfile"

список некоторых используемых пакетов:
github.com/BurntSushi/toml 
github.com/sirupsen/logrus
github.com/gorilla/mux
github.com/stretchr/testify
github.com/lib/pq
github.com/golang-migrate/migrate
github.com/grailbio/base
