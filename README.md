#Simple math service

##How to run
* `git clone https://github.com/YegorMaksymchuk/mathservice.git`
* docker build -t mathservice:latest .
* docker run -d -it -p 6000:6000 --name=mathservice mathservice:latest
 