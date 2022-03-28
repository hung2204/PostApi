# Đề tài:
1. Cài đặt Docker
2. Cài đặt Vernemq broker (docker run -e "DOCKER_VERNEMQ_ACCEPT_EULA=yes" -e "DOCKER_VERNEMQ_ALLOW_ANONYMOUS=on" --name vernemq1 -d vernemq/vernemq)
3. Cài đặt Postgresql dạng docker hoặc trực tiếp (HD cài bằng docker https://hub.docker.com/_/postgres)
4. Viết bằng Golang 1 chương trình giả lập 1 thiết bị cảm biến IOT gửi thông số nhiệt độ vào MQTT broker ngẫu nhiên giá trị với min value là -40 và max value là 120
 Cấu trúc bản tin gửi vào broker như sau: {"temperature":"80"}
 Tên topic: DeviceID/Temperature . Ví dụ: device1/temperature
 Interval 10 giây gửi 1 bản tin
5. Ở golang đồng thời subscribe vào topic được gửi bản tin nhiệt độ tới, lưu trữ vào DB postgresql. Cung cấp API cho phép truy xuất giá trị này
6. Sử dụng bootstrap hoặc tự viết 1 chương trình Frontend ngôn ngữ tự chọn mô phỏng dữ liệu lên biểu đồ đường (line chart). Khuyến khích mô phỏng real time bằng websocket, hoặc có thể gọi api liên tục.

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
