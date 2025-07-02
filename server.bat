chcp 65001
ssh root@95.140.154.127 "rm -r //fc-mobile-telegram-bot"
ssh root@95.140.154.127 "git clone -b dev https://github.com/Looken15/fc-mobile-telegram-bot.git /fc-mobile-telegram-bot"
ssh root@95.140.154.127 "docker stop bot"
ssh root@95.140.154.127 "docker rm bot"
ssh root@95.140.154.127 "docker rmi bot"
ssh root@95.140.154.127 "docker buildx build --no-cache -f //fc-mobile-telegram-bot/Dockerfile -t bot //fc-mobile-telegram-bot/."
ssh root@95.140.154.127 "docker run --name bot -p 8080:8080 -d bot"
pause