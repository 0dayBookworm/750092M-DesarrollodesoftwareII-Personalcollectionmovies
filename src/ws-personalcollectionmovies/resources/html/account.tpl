{{template "header" .}}
<div class="container">
    <!-- Page Header -->
    <div class="page-header">
        <center>
            <h1><strong>PERFIL</strong></h1>
        </center>
    </div>
    <!-- Contenedor de perfiles -->
    <div class="row row-offcanvas row-offcanvas-left">
        <!-- Sidebar -->
        <div class="col-xs-6 col-sm-2 sidebar-offcanvas" role="navigation">
            <div class="text-center">
                <a href="//personalcollectionmovies-alobaton.c9users.io/account/profile">
                    <img id="Avatar" src={{.Avatar}} class="avatar img-circle" alt="avatar">
                </a>
                <h4>{{.Username}}</h4>
                <p>{{.Email}}</p>
            </div>
            <br>
            <ul class="nav">
                <li>
                    {{.AccountControl}}
                </li>
                <hr>
                <li>
                    {{.Reports}}
                </li>
            </ul>
        </div>
        <!-- Content -->
        <div class="col-xs-12 col-sm-10">
            {{.AccountContent}}
        </div>
    </div>
</div>
{{template "footer" .}}