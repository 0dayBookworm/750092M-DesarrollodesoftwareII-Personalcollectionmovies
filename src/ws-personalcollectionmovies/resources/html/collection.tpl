<div class="container">
    <!-- Page header -->
    <div class="page-header">
        <center>
            <h1><strong>{{.Title}}</strong></h1>
        </center>
    </div>
    <!-- Modo de visualización -->
    <div class="well well-sm">
        <!--
        <strong>Vista  </strong>
        -->
        <div class="btn-group">
            <a href="#" id="ViewList" class="btn btn-default btn-sm">
                <span class="glyphicon glyphicon-th-list"></span> Lista
            </a>
            <a href="#" id="ViewGrid" class="btn btn-default btn-sm">
                <span class="glyphicon glyphicon-th"></span> Cuadricula
            </a>
        </div>
    </div>
    <!-- Las peliculas de la busqueda -->
    <div id="Movies" class="row list-group">
        {{str2html .MovieShorts}}
    </div>
    <!-- Paginador -->
    <div class="text-center">
        <ul class="pagination"></ul>
    </div>
</div>