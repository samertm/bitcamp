$(".name").replaceWith(function() {
    return "<td class='name'><a href='http://www.amazon.com/s/ref=nb_sb_noss?field-keywords=" + $(this).text() + "'>" + $(this).text() + "</a></td>";
});