pipeline {
     agent any

     stages {
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