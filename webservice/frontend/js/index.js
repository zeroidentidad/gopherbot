// simple ajax post
$(document).ready(function () {
    let path=window.location.protocol+'//'+window.location.hostname+':'+window.location.port;
    $("form").submit(function (event) {
        let formData={
            description: $("#description").val(),
            level: $("#level").val(),
            challenge_type: $("#challenge_type").val(),
        };

        $.ajax({
            type: "POST",
            url: `${path}/challenges`,
            data: formData,
            dataType: "json",
            encode: true,
        }).done(function (data) {
            // log response
            console.log(data);
        });

        event.preventDefault();
    });
});