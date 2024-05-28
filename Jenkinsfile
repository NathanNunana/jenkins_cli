pipeline{
  agent {
      docker {
        image "gcc:latest"
      }
  } 
  environment {
    DOCKERHUB_CREDENTIALS=credentials("dockerhub")
  }
  stages {
    stage("Checkout"){
        steps {
            checkout scm
        }
    }
    stage("Build"){
      steps {
        script {
          echo "Building image"
          sh "make build:image"
        }
      }
    }
    stage("Docker Login"){
      steps {
        script {
            withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'password', usernameVariable: 'username')]) {
              sh "echo ${password} | docker login -u ${username} --password-stdin"
            }
          }
      }
    }
    stage("Deploy"){
       steps {
        script {
          sh "make deploy"
        }
      } 
    }
  }
}