# Sử dụng Node.js để build ứng dụng
FROM node:18 AS build

# Tạo thư mục làm việc trong container
WORKDIR /app

# Sao chép tệp package.json và package-lock.json
COPY package*.json ./

# Cài đặt dependencies
RUN npm install

# Thêm biến môi trường để khắc phục lỗi OpenSSL
ENV NODE_OPTIONS=--openssl-legacy-provider

# Sao chép toàn bộ mã nguồn
COPY . .

# Build ứng dụng
RUN npm run build

# Bước cuối: Chạy với một web server nhỏ gọn (như nginx)
FROM nginx:alpine

# Sao chép các file build từ bước trước
COPY --from=build /app/build /usr/share/nginx/html

# Mở cổng 80
EXPOSE 80

# Khởi chạy nginx
CMD ["nginx", "-g", "daemon off;"]
