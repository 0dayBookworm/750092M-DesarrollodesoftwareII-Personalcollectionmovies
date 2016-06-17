<ul class="nav navbar-nav navbar-right">
    <li class ="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
            <i class="glyphicon glyphicon-user"></i> 
            <strong>  {{.Username}}  </strong>
            <i class="glyphicon glyphicon-chevron-down"></i> 
        </a>
        <ul class="dropdown-menu">
            <li>
            
                <a href="//personalcollectionmovies-alobaton.c9users.io/profile" id="Profile" name="Profile">
                <!--
                <a href="//personalcollectionmovies-alobaton.c9users.io/profile" id="Profile" name="Profile">
                <a href="#" id="Profile" name="Profile">
                -->
                    <div class="row" id="ProfileUsername" name ="ProfileUsername">
                        <strong>{{.Username}}</strong>
                    </div>
                    <div class="row">
                        <small>Actualizar datos</small>
                    </div>
                </a>
            </li>
            <hr class="divider"></hr>
            <li>
                <a href="//personalcollectionmovies-alobaton.c9users.io/logout" id="Logout" name="Logout">
                    <small> <i class="glyphicon glyphicon-log-out"></i>  Cerrar Sesion </small>
                </a>
            </li>
        </ul>
        
    </li>
</ul>