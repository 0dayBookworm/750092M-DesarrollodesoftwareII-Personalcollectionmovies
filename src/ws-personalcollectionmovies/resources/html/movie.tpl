{{template "header" .}}
<div class="container">
    <!-- Page header -->
    <div class="page-header">
        <center>
            <h1><strong>DETALLES DE PELÍCULA</strong></h1>
        </center>
    </div>
    <div class="container">
        <!-- 
            Titulo de la pelicula 
        <i class="glyphicon glyphicon-bookmark"></i>
        -->
        <h2 class="head">
            {{.Title}}
            <small>
                <strong>
                   ({{.Year}})
                </strong>
            </small>
        </h2>
        <h4>
            {{.OriginalTitle}}
        </h4>
    </div>
    <div class="panel-default">
        <!-- Contenedor: Titulo de pelicula -->
        <!-- Fin titulo de la pelicula -->
        <!-- Información de la pelicula -->
        <div class="panel-body">
            <div class="row">
                <div class="col-md-6">
                    <div class="row">
                        <center class="col-md-4">
                            <img class="img-thumbnail img-responsive" src={{.Poster}}>
                            <h4>CALIFICACIÓN </h4>
                            <input id="Rating" type="hidden" class="rating" />
                            <!-- Controls -->
                            <hr>
                            <div class="btn-group">
                                <button type="button" class="btn dropdown">
                                    <i class="dropdown-toggle glyphicon glyphicon-plus" data-toggle="dropdown"></i>
                                    <ul class="dropdown-menu" aria-labelledby="dLabel">
                                        <li>
                                            <a id="ViewListAdd">
                                                <small> <i class="glyphicon glyphicon-eye-close"></i>  Añadir a Vistas </small>
                                            </a>
                                        </li>
                                        <li>
                                            {{str2html .WatchListContent}}
                                        </li>
                                        <li>
                                            <a href="#">
                                                <small> <i class="glyphicon glyphicon-list"></i>  Mis listas personalizadas</small>
                                            </a>
                                        </li>
                                    </ul>
                                </button>
                                <button type="button" class="btn dropdown">
                                    <i class="dropdown-toggle fa fa-share-alt" data-toggle="dropdown"></i>
                                    <ul class="dropdown-menu" aria-labelledby="dLabel">
                                        <li>
                                            <a href="http://www.facebook.com/sharer.php?u={{.PageUrl}}" target="_blank" onclick="window.open(this.href,'targetWindow','toolbar=no,location=0,status=no,menubar=no,scrollbars=yes,resizable=yes,width=600,height=250'); return false">
                                                <small> <i class="fa fa-facebook fa-2"></i>  Compartir en Facebook</small>
                                            </a>
                                        </li>
                                    </ul>
                                </button>
                            </div>
                            <!-- End controls -->
                        </center>
                        <div class="col-md-8">
                            <h3>FICHA TECNICA </h3>
                            <p>
                                <strong>DURACIÓN: </strong>
                                <br>
                                <strong> {{.Runtime}} </strong>
                                <br> mn
                            </p>
                            <p>
                                <strong>FECHA DE ESTRENO: </strong>
                            </p>
                            <div class="row">
                                <div class="col-md-3">
                                    <div class="meta-date text-center">
                                        <p><span class="date">{{.Released}}{{.Day}}</span><span>{{.Month}}</span><span>{{.Year}}</span></p>
                                    </div>
                                </div>
                            </div>
                            <p><strong>COMPAÑIAS PRODUCTORAS: </strong> {{str2html .ProductionCompanies}}</p>
                            <p><strong>GENEROS: </strong>{{str2html .Genres}}</p>
                            <a href={{.Website}} target="_blank">Pagina oficial <i class="glyphicon glyphicon-new-window"></i></a>
                        </div>
                    </div>
                </div>
                <div class="col-md-6">
                    <h3>TRAILER</h3>
                    <div class="embed-responsive embed-responsive-16by9">
                        <iframe class="embed-responsive-item" src="https://www.youtube.com/embed/{{.Trailer}}" frameborder="0" allowfullscreen></iframe>
                    </div>
                    <br>
                    <p class="justify">
                        <strong>SINOPSIS: </strong>
                        <br> {{.Plot}}
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
                                <div class="col-xs-11">
                                    <textarea type="text" id="Comment" name="Comment" class="form-control" placeholder="¿Que te parecio la pelicula?" row="2" cols="70"></textarea>
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
                        <h3>COMENTARIOS</h3>
                        <!-- Comentarios -->
                        <ul class="media-list">
                            <li class="media">
                                <div class="media-body">
                                    <h4 class="media-heading">Un usuario</h4> Un comentario.
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