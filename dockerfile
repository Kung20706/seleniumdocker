FROM golang:latest
# Adding trusting keys to apt for repositories
# RUN apt-get update -y && \
#     apt-get install -y \
#     wget \
#     curl \
#     unzip

# # 下载并安装 Chrome 浏览器
# RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
#     echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list && \
#     apt-get update -y && \
#     apt-get install -y google-chrome-stable

# # 下载并安装 ChromeDriver
# RUN CHROMEDRIVER_VERSION=$(curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE) && \
#     wget -q -O /tmp/chromedriver.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip && \
#     unzip /tmp/chromedriver.zip -d /project_name && \
#     rm /tmp/chromedriver.zip
WORKDIR /project_name
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o main .
EXPOSE 8087
CMD ["./main"]