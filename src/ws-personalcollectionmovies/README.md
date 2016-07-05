## ws-personalcollectionmovies
Servicio web para una empresa de exhibición y distribución de cine que permite almacenar la información sobre sus gustos en cine.
# Guía para ejecutar la aplicación #
Para ejecutar la aplicación en se debe seguir la siguiente linea de comandos:
1. Accedemos a la ruta en la cual descargamos el repositorio.
```sh
cd <app_path>
```
2. Ejecutamos el comando bee run y nos debera mostrar lo siguiente.
```sh
bee   :1.4.1
beego :1.6.1
Go    :go version go1.6 linux/amd64

2016/06/05 02:44:10 [INFO] Uses 'ws-personalcollectionmovies' as 'appname'
2016/06/05 02:44:10 [INFO] Initializing watcher...
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies/controllers)
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies)
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies/log)
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain)
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies/routers)
2016/06/05 02:44:10 [TRAC] Directory(/home/ubuntu/workspace/src/ws-personalcollectionmovies/util)
2016/06/05 02:44:10 [INFO] Start building...
2016/06/05 02:44:14 [SUCC] Build was successful
2016/06/05 02:44:14 [INFO] Restarting ws-personalcollectionmovies ...
2016/06/05 02:44:14 [INFO] ./ws-personalcollectionmovies is running...
2016/06/05 02:44:14 [asm_amd64.s:1998][I] http server Running on :8080
```
# Guía para generar objetos de dominio #
Como herramienta de mapeo relacional se utilizo xo. Si no conoces la herramienta te recomendamos dirigirte a [knq/xo](https://github.com/knq/xo).
Si la base de datos sufrio alguna modificación y es necesario generar nuevamente los objetos de dominio se deben seguir los siguientes pasos:
1. Debes borrar el contenido de la carpeta /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain el cual contiene los actuales objetos de dominio.
2. Debes ejecutar la siguiente linea de comandos para generar nuevamente los objetos de dominio.
```sh
xo pgsql://sotilstfjbrplf:pr-JQDyUOTCcm8FzQRmpm8a_g0@ec2-54-243-62-211.compute-1.amazonaws.com/d59ne6oopddc7q -o /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain 
```
**NOTE:** Las clases de dominio NO deberan ser modificadas, si necesitamos añadir una funcionalidad adicional debemos crearlas en la ruta /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain/custom. Bajo el siguiente estandar para el nombramiento de este tipoi de archivos nombre_del_ objeto_de_dominio_custom.go

# Algunos métodos JS utiles #

$("html").html("Texto html");

# Kill a Process on specific port #

fuser -k 8080/tcp