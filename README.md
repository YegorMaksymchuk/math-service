#Simple math service

##How to run
* git clone https://github.com/YegorMaksymchuk/mathservice.git
* cd mathservice
* docker compose up --build 
 ###or
 * sh build-multi-stage.sh
 
 ###Make a test
 * curl -X GET http://localhost:6000/math?amount=2
 
 ### You will see something like:
 [
   {
     "question": "7+3=?",
     "answer": "10"
   },
   {
     "question": "4+8=?",
     "answer": "12"
   }
 ]