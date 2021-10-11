# post
MQTTBROKER
   cài đặt:
	sudo apt-get update
	sudo apt-get install mosquitto
   ktra phiên bản:
	sudo systemctl status mosquitto
   publish:
	mosquitto_sub -t "post" -u "mqttbroker" -P "123"
	mosquitto_pub -m "Hello Hung"  -t "post" -u "mqttbroker" -P "123"
   gỡ cài đặt:
	sudo apt-get purge --remove mosquitto*
	sudo apt-add-repository --remove ppa:mosquitto-dev/mosquitto-ppa

POSTGRESQL
	sudo -u postgres psql -c "SELECT version();"
   vào database mydb:
	psql -h 192.168.15.128 -d mydb -U myuser
	(psql -U myuser -h localhost  -d mydb)
   chọn bảng:
	select * from hung;
   reset:
	sudo ufw allow 5432/tcp
	sudo systemctl restart postgresql

CREATE TABLE user (
  devideid SERIAL PRIMARY KEY,
  value INT,
  time timestamp
);
