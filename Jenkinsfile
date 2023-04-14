pipeline {
     agent any

     stages {
         stage('Checkout') {
             steps {
                 checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[credentialsId: 'github-credentials', url: 'git@github.com:noczero/ZeroAPI-go.git']])
             }
         }
         stage('Deploy') {
             steps {
                 echo 'Deploying....'
                 script {
                   // SSH into the server and execute the script
                   sshagent(credentials: ['ssh-remote-zero']) {
                     sh "ssh -o StrictHostKeyChecking=no ${env.ZERO_SERVER_USERNAME}@${env.ZERO_SERVER_HOST} 'cd ~/${env.JOB_NAME} && ./deploy.sh'"
                   }
                 }
             }
         }
     }
}