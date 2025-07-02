chcp 65001
ssh root@95.140.154.127 "rm -r //karama"
scp -r C:\Users\artem\Desktop\fc-mobile-telegram-bot root@95.140.154.127:/karama
ssh root@95.140.154.127 "docker stop bot"
ssh root@95.140.154.127 "docker rm bot"
ssh root@95.140.154.127 "docker rmi bot"
ssh root@95.140.154.127 "docker buildx build --no-cache -f //karama/Dockerfile -t bot //karama/."
ssh root@95.140.154.127 "docker run --name bot -p 8080:8080 -d bot"
pause