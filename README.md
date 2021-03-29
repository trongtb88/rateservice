# rateservice
The service for currency rate management in Golang

## How to run this service at localhost
1. Start your mysql at your localhost machine successfully
2. Git clone this repo
3. Change file .env which is mapped with your config (DB_USER, DB_PASSWORD, DB_HOST), please note commment DB_HOST at docker, uncomment DB_HOST using 127.0.0.1
4. Create database ***rate_db*** by yourself
5.  Go to terminal at root of project
```sh
   go get .    
   go run main.go
```

6. If have some logs at console like, server started and worked successfully

```sh
2021/03/29 00:20:32 Listening to port 8687
2021/03/29 00:20:32 Reading xml file
2021/03/29 00:20:34 Parsing done
2021/03/29 00:20:34 Found 2048 curreny rates 
2021/03/29 00:20:34 Load data from url into database successfully
```

7. Go to your browser like Chrome, type URL :  http://localhost:8687/rates/latest, check response is returned


## How to run this service at docker
1. Run DOCKER DEAMON at your machine successfully
2. Make sure don't have any image mysql is running at port 3306, otherwise you will have 1 error
3. Make sure don't have any image is running at port 8687, otherwise you will have 1 error
2. Git clone this repo
3. Change values config at file .env which is mapped with your config (DB_USER, DB_PASSWORD, DB_HOST), please note uncommment DB_HOST at docker, comment using 127.0.0.1, and value of TEST_DB_HOST=rate-mysql must save value at docker-composer.yml
5.  Go to terminal at root of project
```sh
   docker-composer up --build 
```

6. If have some logs at console like, server started and worked successfully

```sh
rate_db_mysql | 2021-03-29T00:16:59.979485Z 0 [Note]   - '::' resolves to '::';
rate_db_mysql | 2021-03-29T00:16:59.979567Z 0 [Note] Server socket created on IP: '::'.
rate_db_mysql | 2021-03-29T00:17:00.041741Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
rate_db_mysql | 2021-03-29T00:17:00.048674Z 0 [Warning] 'user' entry 'root@rate-mysql' ignored in --skip-name-resolve mode.
rate_db_mysql | 2021-03-29T00:17:00.093867Z 0 [Note] Event Scheduler: Loaded 0 events
rate_db_mysql | 2021-03-29T00:17:00.094413Z 0 [Note] mysqld: ready for connections.
rate_db_mysql | Version: '5.7.33'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
rate_app      | We are connected to the mysql database
rate_app      | 2021/03/29 00:20:32 /app/api/db/config.go:34
rate_app      | [1.480ms] [rows:-] SELECT DATABASE()
rate_app      | 
rate_app      | 2021/03/29 00:20:32 /app/api/db/config.go:34
rate_app      | [5.688ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'rate_db' AND table_name = 'currency_rates' AND table_type = 'BASE TABLE'
rate_app      | 
rate_app      | 2021/03/29 00:20:32 /app/api/db/config.go:34
rate_app      | [19.524ms] [rows:0] CREATE TABLE `currency_rates` (`date` datetime(3) NOT NULL,`currency_code` varchar(5) NOT NULL,`rate` decimal(15,6) NOT NULL,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,PRIMARY KEY (`date`,`currency_code`))
rate_app      | 2021/03/29 00:20:32 Listening to port 8687
rate_app      | 2021/03/29 00:20:32 Reading xml file
rate_app      | 2021/03/29 00:20:34 Parsing done
rate_app      | 2021/03/29 00:20:34 Found 2048 curreny rates 
rate_app      | 2021/03/29 00:20:34 Load data from url into database successfully.
rate_app      | 2021/03/29 00:20:52 200

```

7. Go to your browser like Chrome, type URL :  http://localhost:8687/rates/latest, check response is returned

## How to test this service at localhost

1. Start your mysql at your localhost machine successfully
2. Git clone this repo
3. Change values config at .env which is mapped with your config (TEST_DB_USER, TEST_DB_PASSWORD, TEST_DB_HOST), please note commment DB_HOST at docker, uncomment using 127.0.0.1
4. Create database ***test_rate_db*** by yourself
5.  Go to terminal at root of project
```sh
   cd tests
   go test
```
6. If you see some logs like, mean tests passed.

```sh
2021/03/29 07:38:13 /Users/trtran/go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:165
[2.280ms] [rows:0] DROP TABLE IF EXISTS `currency_rates` CASCADE

2021/03/29 07:38:13 /Users/trtran/go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:170
[0.081ms] [rows:0] SET FOREIGN_KEY_CHECKS = 1;

2021/03/29 07:38:13 /Users/trtran/SProjects/rateservice/tests/setup_test.go:55
[0.094ms] [rows:-] SELECT DATABASE()

2021/03/29 07:38:13 /Users/trtran/SProjects/rateservice/tests/setup_test.go:55
[0.765ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test_rate_db' AND table_name = 'currency_rates' AND table_type = 'BASE TABLE'

2021/03/29 07:38:13 /Users/trtran/SProjects/rateservice/tests/setup_test.go:55
[4.056ms] [rows:0] CREATE TABLE `currency_rates` (`date` datetime(3) NOT NULL,`currency_code` varchar(5) NOT NULL,`rate` decimal(15,6) NOT NULL,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,PRIMARY KEY (`date`,`currency_code`))
2021/03/29 07:38:13 Successfully refreshed table
2021/03/29 07:38:13 200
PASS
ok      github.com/trongtb88/rateservice/tests  0.312s

```







