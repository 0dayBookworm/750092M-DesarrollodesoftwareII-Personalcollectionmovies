{{template "header" .}}
<div class="container">
    <div class="page-header">
        <center>
            <h1><strong>FORMULARIO DE REGISTRO</strong></h1></center>
    </div>
    <!-- Formulario de registro -->
    <!--
    <form class="form-horizontal" id="Useraccount" name="Useraccount"  action="https://personalcollectionmovies-alobaton.c9users.io/register/create" method="post">
    -->
    <form class="form-horizontal" id="RegisterForm" name="RegisterForm" data-toggle="bootstrapValidator">
        <!-- Username and mail-->
        <div class="form-group">
            <label class="control-label col-xs-3">Nombre de usuario:</label>
            <div class="col-xs-9">
                <input type="text" id="Username" name="Username" class="form-control text-lowercase" placeholder="Nombre de usuario">
            </div>
        </div>
        <div class="form-group">
            <label class="control-label col-xs-3">Email:</label>
            <div class="col-xs-9">
                <input type="email" id="Email" name="Email" class="form-control text-lowercase" placeholder="Email">
            </div>
        </div>
        <hr class="container">
        <!-- Fin username and mail -->
        <!-- Contraseña y confirmación-->
        <div class="form-group">
            <label class="control-label col-xs-3">Contraseña:</label>
            <div class="col-xs-9">
                <input type="password" id="Password" name="Password" class="form-control" placeholder="Contraseña">
            </div>
        </div>
        <div class="form-group">
            <label class="control-label col-xs-3">Confirmar Contraseña:</label>
            <div class="col-xs-9">
                <input type="password" id="ConfirmPassword" name="ConfirmPassword" class="form-control" placeholder="Confirmar Contraseña">
            </div>
        </div>
        <small><i>La contraseña debe contener almenos una mayuscula y un numero.</i></small>
        <hr class="container">
        <!-- Fin contraseña y confirmación -->
        <!-- Primer nombre -->
        <div class="form-group">
            <label class="control-label col-xs-3">Primer nombre:</label>
            <div class="col-xs-9">
                <input type="text" id="FirstName" name="FirstName" class="form-control" placeholder="Primer Nombre">
            </div>
        </div>
        <!-- Segundo nombre -->
        <div class="form-group">
            <label class="control-label col-xs-3">Segundo nombre:</label>
            <div class="col-xs-9">
                <input type="text" id="SecondName" name="SecondName" class="form-control" placeholder="Segundo Nombre">
            </div>
        </div>
        <!-- Apellidos -->
        <div class="form-group">
            <label class="control-label col-xs-3">Apellidos:</label>
            <div class="col-xs-9">
                <input type="text" id="LastName" name="LastName"  class="form-control" placeholder="Apellido">
            </div>
        </div>
        <!-- Fecha de nacimiento -->
        <div class="form-group">
            <label class="control-label col-xs-3">F. Nacimiento:</label>
            <div class="col-xs-9">
                <div class="input-group date" id="Date">
                    <span class="input-group-addon"><i class="glyphicon glyphicon-calendar"></i></span>
                    <input type="text" id="BirthDate" name="BirthDate" class="form-control" placeholder="yyyy-MM-dd">
                </div>
            </div>
        </div>
        <!-- Genero -->
        <div class="form-group">
            <label class="control-label col-xs-3">Genero:</label>
            <div class="col-xs-9">
                <div class="col-xs-9">
                    <label class="radio-inline">
                        <input type="radio" id="Gender" name="Gender" value="male"> Masculino
                    </label>
                </div>
                <div class="col-xs-9">
                    <label class="radio-inline">
                        <input type="radio" id="Gender" name="Gender" value="female"> Femenino
                    </label>
                </div>
            </div>
        </div>
        <!-- Fin genero -->
        <hr class="container">
        <!-- Terminos y condiciones -->
        <div class="form-group">
            <div class="col-xs-offset-3 col-xs-9">
                <label class="checkbox-inline">
                    <input type="checkbox" id="TermsAndConditions" name="TermsAndConditions" value="agree"> Accepto <a href="#">Terminos y condiciones</a>.
                </label>
            </div>
        </div>
        <br>
        <!-- Controles -->
        <div class="form-group">
            <div class="col-xs-offset-3 col-xs-9">
                <!--
                <button type="submit" class="btn btn-primary" aria-label="Left Align">
                -->
                <button type="submit" id="CreateUseraccount" class="btn btn-primary" aria-label="Left Align" disabled>
                    <i class="glyphicon glyphicon-ok"></i>
                    Registrarse
                </button>
                <button type="reset" class="btn btn-default">Limpiar</button>
            </div>
        </div>
    </form>
</div>
{{template "footer" .}}