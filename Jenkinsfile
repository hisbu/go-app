pipeline {
  agent any

  environment {
    // CI = true
    DOCKER_TAG = getDockerTag()
  }

  stages {
   //stage empat
    stage ("build docker images"){
      steps{
        script {
          app = docker.build("hisbu/my-goapp")
        }
      }
    }

    //stage lima
    stage ("test docker image"){
      steps{
        sh 'docker run -d --rm --name testimage -p 9000:9000 hisbu/my-goapp'
        // input message: "Finished test image? (Click procced to continue)"
      }
    }

    //stage enam
    stage("clean up docker test"){
      steps{
        sh 'docker stop testimage'
      }
    }

    //stage tujuh
    stage ("push image to registry"){
      steps{
        script{
          docker.withRegistry('https://registry.hub.docker.com', 'dockerhub-hisbu'){
            app.push("${DOCKER_TAG}")
            app.push("latest")
          }
        }
      }
    }

    // stage delapan
    stage("cleanup docker images"){
      steps{
        sh 'docker rmi hisbu/my-goapp'
        // sh "docker rmi registry.hub.docker.com/hisbu/reactapp-jcde:'${DOCKER_TAG}'"
      }
    }

    //stage sembilan
    // stage("deploy app to kubernetes cluster"){
    //   steps{
    //     sh "chmod +x changeTag.sh"
    //     sh "./changeTag.sh ${DOCKER_TAG}"
    //     withKubeConfig([credentialsId: 'kubecconfig-clusterjcde', serverUrl:'https://34.101.207.52']){
    //       sh 'kubectl apply -f reactapp-config.k8s.yaml'
    //     }
    //   }
    // }

  }
}

def getDockerTag(){
  def tag = sh script: "git rev-parse HEAD", returnStdout: true
  return tag
}