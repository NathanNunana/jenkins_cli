pipeline{
  agent any 
  environment {
    DOCKERHUB_CREDENTIALS=credentials("dockerhub")
    image='jcli'
    username='ghost023'
    version='latest'
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
          sh "docker build -t ${username}/${image}:${version} ."
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
          echo "Deploying image to registry"
          docker push ${username}/${image}:${version}
        }
      } 
    }
  }
}
