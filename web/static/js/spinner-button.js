$('#load').on('click', function() {
    var $this = $(this);
    $this.button('searching');
    $this.parents('form#search-form').submit();
    setTimeout(function() {
        $this.button('reset');
    }, 59000);
});
