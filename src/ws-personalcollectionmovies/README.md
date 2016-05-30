## ws-personalcollectionmovies
Servicio web para una empresa de exhibición y distribución de cine que permite almacenar la información sobre sus gustos en cine.

# Guía para generar objetos de dominio #

Como herramienta de mapeo relacional se utilizo xo. Si no conoces la herramienta te recomendamos dirigirte a [knq/xo](https://github.com/knq/xo).

Si la base de datos sufrio alguna modificación y es necesario generar nuevamente los objetos de dominio se deben seguir los siguientes pasos:

1. Debes borrar el contenido de la carpeta /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain el cual contiene los actuales objetos de dominio.
2. Debes ejecutar la siguiente linea de comandos para generar nuevamente los objetos de dominio.

```sh
xo pgsql://sotilstfjbrplf:pr-JQDyUOTCcm8FzQRmpm8a_g0@ec2-54-243-62-211.compute-1.amazonaws.com/d59ne6oopddc7q -o /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain 
```

**NOTE:** Las clases de dominio NO deberan ser modificadas, si necesitamos añadir una funcionalidad adicional debemos crearlas en la ruta /home/ubuntu/workspace/src/ws-personalcollectionmovies/model/domain/custom. Bajo el siguiente estandar para el nombramiento de este tipoi de archivos nombre_del_ objeto_de_dominio_custom.go
