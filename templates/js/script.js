$(".name").replaceWith(function() {
    return "<td class='name'><a href='http://www.amazon.com/s/ref=nb_sb_noss?field-keywords=" + $(this).text() + "'>" + $(this).text() + "</a></td>";
});

$(".price").text(function() {
    if ($(this).text().length - $(this).text().indexOf(".") > 3) {
        return $(this).text().substring($(this).text().indexOf("."), $(this).text().indexOf(".") + 2);
    } else {
        return $(this).text()
    }
});

$(".macro").text(function() {
    if ($(this).text().length > 3) {
        return $(this).text().substring(0,$(this).text().length - 3) + " g";
    } else {
        return $(this).text() + " mg";
    }
});