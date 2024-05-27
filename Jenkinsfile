pipeline{
  agent: { dockerfile true }
  stages {
    stage("Build"){
      script {
        sh "make build:image"
      }
    }
    stage("Deploy"){
       script {
          sh "make deploy"
        } 
    }
  }
}
