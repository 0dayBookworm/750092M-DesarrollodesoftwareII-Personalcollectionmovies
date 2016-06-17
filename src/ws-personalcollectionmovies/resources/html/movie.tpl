{{template "header" .}}

<div class="container">
    <!-- Page header -->
    <div class="page-header">
        <center>
            <h1><strong>DETALLES DE PELÍCULA</strong></h1>
        </center>
    </div>
  
    <div class="panel panel-default">
        <!-- Contenedor: Titulo de pelicula -->
        <div class="panel-heading">
            <!-- Titulo de la pelicula -->
            <h1>TITULO: </h1>
            <p class="panel-title pull-left">Capitan America: Civil War</p>
            <!-- Controls -->
            <div class="row pull-right">
                <div class ="col-md-1 dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <i class="glyphicon glyphicon-plus"></i>
                    </a>
                    <ul class="dropdown-menu">
                        <li>
                            <a href="#">
                                <small> <i class="glyphicon glyphicon-eye-open"></i>  Añadir a vistas </small>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <small> <i class=""></i>  Añadir a películas por ver </small>
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <small> <i class=""></i>  Mis listas </small>
                            </a>
                        </li>
                    </ul>
                </div>   
                <i class="col-md-1 glyphicon glyphicon-share-alt"></i>
            </div>
            <!-- Fix zone -->
            <div class="clearfix"></div>
        </div>
        <!-- Fin titulo de la pelicula -->
        
        
        <!-- Información de la pelicula -->
        <div class="panel-body">
            <div class="row">
                <div class="col-md-4">
                    <center>
                        <img src="//personalcollectionmovies-alobaton.c9users.io/public/images/CivilWar.jpg">
                        <h1>CALIFICACIÓN </h1>
                        <input id="Rating" type="hidden" class="rating"/>
                    </center>
                </div>
                <div class="col-md-8">
                    <div class=row>
                        
                        <div class="col-md-4">
                            <h1>FICHA TECNICA </h1>
                            <p class=""><strong>País: </strong> USA, Germany  </p>
                            <p class=""><strong>Año: </strong>2016</p>
                            <p class=""><strong>Categoria: </strong>  Action, Adventure </p>
                            <p class=""><strong>Fecha de lanzamiento: </strong> 06 May 2016 </p>
                            <p class=""><strong>Director: </strong> Anthony Russo, Joe Russo</p>
                            <p class="">
                                <strong>Actores: </strong>
                                Chris Evans, Robert Downey Jr., Scarlett Johansson, Sebastian Stan
                            </p>
                            <p class=""><strong>Edad minima: </strong></p> 
                            <p><strong>Premios: </strong>N/A</p>
                        </div>
                        <div class="col-md-8">
                            <h1>TRAILER</h1>
                            <div class="embed-responsive embed-responsive-16by9">
                                <iframe class="embed-responsive-item" src="https://www.youtube.com/embed/dKrVegVI0Us" frameborder="0" allowfullscreen></iframe>
                            </div>
                        </div>
                    </div>
                        <p>
                            <strong>Descripción de la pelicula: </strong>
                            Political interference in the Avengers' activities causes a rift between former allies Captain America and Iron Man.
                        </p>
                </div>
            </div>
        </div>
        <!-- Fin informacion de la pelicula -->
        <!-- Seccion de comentarios. -->
    
        <div class="panel-footer">
            
            <!-- Realizar un comentario -->
            <div class="container">
                <p id="CommentSection" class="dropdown-toggle" data-toggle="collapse" data-target="#CommentCollapse">
                    <i class="glyphicon glyphicon-comment"></i>
                    <strong>  Comentar  </strong>
                    <i class="glyphicon glyphicon-chevron-down"></i>
                </p>
                <div id="CommentCollapse" class="collapse">
                    <div class="container">
                        <!-- Formulario de comentario -->
                        <form class="form-horizontal" id="CommentForm" name="CommentForm">
                            <div class="form-group">
                                <div class="col-xs-9">
                                    <textarea type="text" id="Comment" name="Comment" class="form-control" placeholder="¿Que te parecio la pelicula?" row="2" cols="70" ></textarea>
                                </div>
                            </div>
                            <div class="form-group">
                                <div class="col-xs-3">
                                    <button type="button" id="CommentButton" class="btn btn-primary" aria-label="Left Align">
                                        <i class="glyphicon glyphicon-ok"></i>
                                        Comentar
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                    <div class="container">
                        <h1>COMENTARIOS</h1>
                        <!-- Comentarios -->
                        <ul>
                            <li class="col-xs-9">
                                <div class="panel ">
                                    <div class ="panel-heading comment-header">
                                        <strong>
                                            Otro usuario
                                        </strong>
                                        <p>
                                            <small>
                                            Hace 10 minutos.
                                            </small>
                                        </p>
                                    </div>
                                    <div class="panel-body">
                                        <p>Un comentario.</p>
                                    </div>
                                </div>
                            </li>
                            
                            <li class="col-xs-9">
                                <div class="panel ">
                                    <div class ="panel-heading comment-header">
                                        <strong>
                                            Otro usuario
                                        </strong>
                                        <p>
                                            <small>
                                            Hace 10 minutos.
                                            </small>
                                        </p>
                                    </div>
                                    <div class="panel-body">
                                        <p>Otro comentario.</p>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        <!-- Fin sección de comentarios. -->
    </div>



</div>

{{template "footer" .}}