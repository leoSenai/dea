# dea - branch beta

<p>This branch is the <b>homologation</b> branch</p>
<ul>
  <li>Version: 0.0.1</li>
  <li>Authors: 
    <i>
      <a href="https://github.com/davidsidor" target="_blank">David Sidro</a>, 
      <a href="https://github.com/daniloserafim" target="_blank">Danilo de Matos Serafim</a>, 
      <a href="https://github.com/dariokrugerjunior" target="_blank">Dario Kruger Junior</a>, 
      <a href="https://github.com/FelipeJanotte" target="_blank">Felipe Janotte</a>, 
      <a href="https://github.com/leoSenai" target="_blank">Leonardo</a>, 
      <a href="https://github.com/RodrigoPerozin" target="_blank">Rodrigo Perozin</a>
    </i>
  </li>
</ul>

# How to setup and run backend

<ol>
  <li>Clone this git repository</li>
  <li>Checkout the 'beta' branch</li>
  <li>Install <a href='https://dev.mysql.com/get/Downloads/MySQLInstaller/mysql-installer-web-community-8.0.34.0.msi'><i>MySQL Community 8.0</i></a> and all packages contained in</li>
  <li>Configure the MySQL and the <i>MySQL Workbench</i></li>
  <li>Open the <i>MySQL Workbench</i>, go inside a instance and open the archive .sql contained on <a href='https://drive.google.com/drive/u/0/'>Google Drive</a></li>
  <li>Run the file and it's almost done..</li>
  <li>Make sure that the instance is on localhost:3306 and both the user and password is set to 'root'</li>
  <li>install <i><a href='https://go.dev/dl/go1.21.0.windows-amd64.msi'>Golang</a></i> on your machine</li>
  <li>Go to a terminal and go to directory '/Backend'</li>
  <li>run command 'go run main.go'</li>
</ol>

# How to setup and run frontend

<ol>
  <li>Clone this git repository</li>
  <li>Checkout the 'beta' branch</li>
  <li>Install <a href="https://nodejs.org/pt-br/download">Node JS</a></li>
  <li>Open a terminal and type:
    <pre>
      <code>$ cd /dea/FrontEnd -> To change for project directory </code>
      <code>$ npm i -> To install project dependencies</code>
      <code>$ npm run serve -> To run project</code>
    </pre>
  </li>
  <li>Open your favorite navitor and go to <a href="http://localhost:5173/">http://localhost:5173/</a></li>
</ol>
