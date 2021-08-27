// simple ajax post
$(document).ready(function () {
    let path=window.location.protocol+'//'+window.location.hostname+':'+window.location.port;
    $("form").submit(function (e) {
        let description=$("#description").val();
        let level=$("#level").val();
        let challengetype=$("#challengetype").val();

        if (description!=""&&level!=""&&challengetype!="") {
            let formData={
                description,
                level,
                challengetype
            };

            $.ajax({
                type: "POST",
                url: `${path}/challenges`,
                data: formData,
                dataType: "json",
                encode: true,
            }).done(function (data) {
                // log response
                console.log("Response ok:", data);
                alert("Desafio recibido y guardado correctamente 😎");
            });
        } else {
            alert("No se estan mandando valores válidos")
        }

        e.preventDefault();
    });
});