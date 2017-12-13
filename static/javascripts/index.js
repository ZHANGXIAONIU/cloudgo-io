$(function() {
    $("button[name='submit']").click(function() {
        $.post(
            "/",
            {
                name: $("input[name='name']").val(),
                birthday: $("input[name='birth']").val()
            },
            function(data, status) {
                if (status == "success") {
                    $("tbody").append(
                        "<tr>" +
                        "  <td>" + data.name + "</td>" +
                        "  <td>" + data.birthday + "</td>" +
                        "</tr>"
                    );
                }
            }
        );
        return false;
    });
});
