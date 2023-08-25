# 生成私钥:加密/解密需要用到
openssl genrsa -out server.key 2048

# 生成证书：由证书颁发机构CA签名后的证书
openssl req -new -x509 -key server.key -out server.crt -days 36500

# 生成csr: 证书签名请求文件，用于提交给证书颁发机构对证签名
openssl req -new -key server.key -out server.csr

# 通过RSA算法生成证书私钥
openssl genpkey -algorithm RSA -out test.key

#
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=jack" -config openssl.cnf -extensions v3_req

#
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

