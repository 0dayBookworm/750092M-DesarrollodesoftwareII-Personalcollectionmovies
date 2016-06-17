{{template "header" .}}
<div class="container">
    <div class="page-header">
        <center>
            <h1><strong>DATOS DE USUARIO</strong></h1>
        </center>
    </div>
    
    <!--Datos de usuario -->
    <form class="form-horizontal " id="ProfileForm" name="ProfileForm" data-toggle="bootstrapValidator">
        <!-- Username and mail-->
        <div class="form-group">
            <label class="control-label col-xs-3">Nombre de usuario:</label>
            <div class="col-xs-9">
                <input type="text" id="Username" name="Username" class="form-control text-lowercase" value={{.Username}} readonly>
            </div>
        </div>
        <div class="form-group">
            <label class="control-label col-xs-3">Email:</label>
            <div class="col-xs-9">
                <input type="email" id="Email" name="Email" class="form-control text-lowercase" value={{.Email}}>
            </div>
        </div>
        <hr class="container">
        <!-- Fin username and mail -->
        <!-- Primer nombre -->
        <div class="form-group">
            <label class="control-label col-xs-3">Primer nombre:</label>
            <div class="col-xs-9">
                <input type="text" if="FirstName" name="FirstName" class="form-control" value={{.FirstName}}>
            </div>
        </div>
        <!-- Segundo nombre -->
        <div class="form-group">
            <label class="control-label col-xs-3">Segundo nombre:</label>
            <div class="col-xs-9">
                <input type="text" id="SecondName" name="SecondName" class="form-control" value={{.SecondName}}>
            </div>
        </div>
        <!-- Apellidos -->
        <div class="form-group">
            <label class="control-label col-xs-3">Apellidos:</label>
            <div class="col-xs-9">
                <input type="text" id="LastName" name="LastName" class="form-control" value={{.LastName}}>
            </div>
        </div>
        <!-- Fecha de nacimiento -->
        <div class="form-group">
            <label class="control-label col-xs-3">F. Nacimiento:</label>
            <div class="col-xs-9">
                <div class="input-group date" id="Date">
                    <span class="input-group-addon">
                            <span class="glyphicon glyphicon-calendar"></span>
                    </span>
                    <input type="text" id="BirthDate" name="BirthDate" class="form-control" value={{.BirthDate}}>
                </div>
            </div>
        </div>
        <!-- Genero -->
        <div class="form-group">
            <label class="control-label col-xs-3">Genero:</label>
            <div class="col-xs-9">
                <div class="col-xs-9">
                    <label class="radio-inline">
                        <input type="radio"  id="Gender" name="Gender"  value="male" {{.Male}}> Masculino
                    </label>
                </div>
                <div class="col-xs-9">
                    <label class="radio-inline">
                        <input type="radio"  id="Gender" name="Gender"  value="female" {{.Female}}> Femenino
                    </label>
                </div>
            </div>
        </div>
        <!-- Fin genero -->
        <!-- Guardar cambios -->
        <br>
        <div class="form-group">
            <div class="col-xs-offset-3 col-xs-9">
                <button type="submit" id="UpdateProfile" class="btn btn-primary" aria-label="Left Align" disabled>
                    <i class="glyphicon glyphicon-pencil"></i>
                    Guardar Cambios
                </button>
            </div>
        </div>
    </form>


    <hr class="container">


    <!-- Contraseña y confirmación-->
    <form class="form-horizontal" id="ChangePasswordForm" name="ChangePasswordForm">
        <div class="form-group">
            <label class="control-label col-xs-3">Contraseña actual:</label>
            <div class="col-xs-9">
                <input type="password" id="Password" name="Password" class="form-control" id="inputPassword" placeholder="Contraseña actual">
            </div>
        </div>
        <div class="form-group">
            <label class="control-label col-xs-3">Contraseña nueva:</label>
            <div class="col-xs-9">
                <input type="password" id="NewPassword" name="NewPassword" class="form-control" placeholder="Contraseña nueva">
            </div>
        </div>
        <div class="form-group">
            <label class="control-label col-xs-3">Confirmar contraseña:</label>
            <div class="col-xs-9">
                <input type="password" id="ConfirmPassword" name="ConfirmPassword" class="form-control" placeholder="Confirmar Contraseña">
            </div>
        </div>
        <div class="container">
            <small><i>La contraseña debe contener almenos una mayuscula y un numero.</i></small>
        </div>
        <br>
        <div class="form-group">
            <div class="col-xs-offset-3 col-xs-9">
                <button type="button" id="ChangePassword" name="ChangePassword" class="btn btn-primary" aria-label="Left Align">
                    <i class="glyphicon glyphicon-lock"></i>
                    Actualizar Contraseña
                </button>
            </div>
        </div>
    </form>


    <hr class="container">
    <!-- Fin contraseña y confirmación -->
    <!-- Eliminar cuenta -->
    <div class="panel panel-default">
        <div class="panel-heading"><strong>Eliminar tu cuenta </strong><i class="glyphicon glyphicon-alert"></i></div>
        <div class="panel-body">
            <div class="panel">
                <small>Si crees que no volveras a usar tu cuenta y la quieres eliminar, podemos ayudarte a hacerlo. Ten en cuenta que no podras volver a activarla ni recuperar ningun dato. Si quieres que se elimine tu cuenta haz clic en el boton Eliminar Cuenta.</small>
            </div>
            <button type="button" class="btn btn-danger" aria-label="Left Align">
                    <i class="glyphicon glyphicon-remove"></i>
                    Eliminar Cuenta
                </button>
        </div>

    </div>
</div>

{{template "footer" .}}