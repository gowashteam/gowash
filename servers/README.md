# Progress and Documentation for Okey Onyia
## Tasks for week One
-[x]  Comprehensive research on Implementation of Idea and Technology to achive this
-[x] Build a Contineous Integration Pipeline with  Git
-   [x] Alert Team members to test the Integration Pipepline and Mark their names to indicate successful Integration
-   ### Instructions on how to Test Integration
-   ***Step1***
         Clone the repo on your command line using Git with the following command
             **$** `git clone https://github.com/gowashteam/gowash.git`
-   ***Step2***
         Open the file ***testaccess.txt***
         Follow the instruction on the file
-   ***Step3***
         Commit the change then Push the commit back to the remote repo
-   ***Step4***
         Go to https://github.com/gowashteam/gowash/blob/master/testaccess.txt and check to see if the change reflects.
         If you see the change then your Contineous Integration Test Passed. IF not, Please let me know and I will help             figure out the problem.
-   ***Step5*** Check your name below when DONE and then commit and push back to repo
       - [x] @gowashteam/okeyonyia123
      - [ ] @gowashteam/KRUTI
      - [ ] @gowashteam/charmishah928
      - [ ] @gowashteam/Keyaa12
## Task for Week Two
### Get A Virtual Private Server and deploy MongoDB database
-  [ ]  Provision a Virtual Private Server with Digital Ocean
-   [ ] Set the Envirunment variable from your Development Machine with `Docker-Machine Eval`
-   [ ] Spin up a MongoDB Image from DockerHub and run a new instance in a Docker container
-   [ ] Deploy the Container on the Virtual Private Server then Get the Socket for communication with Server pages

###  Build the Server that handles user request to SIGNUP new account
  ### Server Folder Structure
-  Servers
    -   Signup
        - handlers
          -  [ ] ***signup.go***
        - utilities ( ***Re-usable class*** )
           -    [ ] ***sha256.go***
           -    [ ] ***uuid.go***
        - [ ] ***SignupServer.go***

### Build a server that handles user request to ACTIVATE new account
-  Servers
    -   ActivaeUSer
        - handlers
          -   [ ] ***activateUser.go***
        - utilities ( ***Re-usable class*** )
           -    [ ] ***sha256.go***
           -    [ ] ***uuid.go***
        - [ ] ***ActivateUserServer.go***

### Build a server that handles user request to LOGIN user
-  Servers
    -   Login
        - handlers
          -   [ ] ***login.go***
        - utilities ( ***Re-usable class*** )
           -    [ ] ***sha256.go***
           -    [ ] ***uuid.go***
        - [ ] ***loginServer.go***

### [ ] Build a server that handles user request to RECOVER-PASSWORD
-  Servers
    -   RecoverPassword
        - handlers
          -   [ ] ***passwordRecovery.go***
        - utilities ( ***Re-usable class*** )
           -   [ ] ***sha256.go***
           -    [ ] ***uuid.go***
        - [ ] ***PRServer.go***
