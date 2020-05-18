pipeline {
  agent {
    kubernetes {
      yamlFile "pod.yaml"
    }
  }

  stages {
    stage("Test") {
      steps {
        container("golang") {
          sh "go run main.go"
        }
      }
    }
  }
}
