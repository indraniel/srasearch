$('#load').on('click', function() {
    var $this = $(this);
    $this.parents('form#search-form').submit();
});

$('#search-form').bind('submit', function() {
    var $this = $(this);
    $b = $this.find("button#load");
    $b.button('searching');
    setTimeout(function() {
        $b.button('reset');
    }, 59000);
});
