<div class="page-header">
    <center>
        <h2><strong>SEGURIDAD</strong></h2>
    </center>
</div>
<form class="form-horizontal" id="ChangePasswordForm" name="ChangePasswordForm">
    <div class="form-group">
        <label class="control-label col-md-3">Contraseña actual:</label>
        <div class="col-md-9">
            <input type="password" id="Password" name="Password" class="form-control" id="inputPassword" placeholder="Contraseña actual">
        </div>
    </div>
    <div class="form-group">
        <label class="control-label col-md-3">Contraseña nueva:</label>
        <div class="col-md-9">
            <input type="password" id="NewPassword" name="NewPassword" class="form-control" placeholder="Contraseña nueva">
        </div>
    </div>
    <div class="form-group">
        <label class="control-label col-md-3">Confirmar contraseña:</label>
        <div class="col-md-9">
            <input type="password" id="ConfirmPassword" name="ConfirmPassword" class="form-control" placeholder="Confirmar Contraseña">
        </div>
    </div>
    <div class="container">
        <small><i>La contraseña debe tener una longitud de ocho caracteres o más.</i></small>
    </div>
    <br>
    <div class="form-group">
        <div class="col-md-offset-3 col-md-9">
            <button type="button" id="ChangePassword" name="ChangePassword" class="btn btn-primary" aria-label="Left Align">
                    <i class="glyphicon glyphicon-lock"></i>
                    Actualizar Contraseña
                </button>
        </div>
    </div>
</form>
{{if .EnableSecutiry}}
<hr>
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
{{end}}