pipeline{
  agent { dockerfile true }
  stages {
    stage("Build"){
      steps {
        script {
          sh "make build:image"
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
