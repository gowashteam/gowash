function theme_script(){
    var item_visible,
        item_start,
        item_animation,
        item_responsive;
    $("body").width() > 767 ? item_visible = 3 : item_visible = 1;
    $("body").width() > 767 ? item_start = 0 : item_start = -1;
    $("body").width() > 767 ? item_animation = 1 : item_animation = 0;
    $("body").width() > 767 ? item_responsive = false : item_responsive = true;

    "use strict";

    /* Touch Events Isset */
    function is_touch_device() {  
      try {  
        document.createEvent("TouchEvent");  
        return true;  
      } catch (e) {  
        return false;  
      }  
    }
    /* End Touch Events Isset */
    
    /* Beners Slider */
    if ($("#sidebar-baners")[0]) {
        $('#sidebar-baners').carouFredSel({
            height: 'auto',
            item: 1,
            scrol: 1,
            swipe: {
                onMouse: true,
                onTouch: true
            },
            pagination: {
                container: "#sidebar-baners-pager"
            }
        });
        var sidebar_widht = $("body").width() < 979 ? 175 : 228;
        var baners_item_width = sidebar_widht / $("#sidebar-baners-pager > a").length -1;
        $("#sidebar-baners-pager > a").css("width", baners_item_width);
        $("#sidebar-baners-pager > a:first-child").css("width", baners_item_width - 1);
    }
    /* End Beners Slider */

    /* Bener New Slider */
    if ($("#baner-new")[0]) {
    $('#baner-new').carouFredSel({
        height: 'auto',
        item: 1,
        scrol: 1,
        auto: false,
        swipe: {
            onMouse: true,
            onTouch: true
        },
        pagination: {
            container: "#baner-new-pager"
        },
        prev: ".baner-new-prev",
        next: ".baner-new-next"
    });
    var sidebar_widht = $("body").width() < 979 ? 175 : 228;
    var baner_item_width = sidebar_widht / $("#baner-new-pager > a").length -1;
    $("#baner-new-pager > a").css("width", baner_item_width);
    $("#baner-new-pager > a:first-child").css("width", baner_item_width - 1);
    }
    /* End Bener New Slider */
    
    /* Header Two Slider */

    $(".header-two-box-prev").click(function(){
        $("#header-two").trigger('prev')
    });
    $(".header-two-box-next").click(function(){
        $("#header-two").trigger('next')
    });

    if ($("#header-two")[0]) {
        if ($("#header-two li").length < 4) {
            $($("#header-two li")).each(function(){
                $(this).clone().appendTo("#header-two");
            });
        }
        $('#header-two').carouFredSel({
            responsive: item_responsive,
            width: '100%',
            height: 'auto',
            auto: true,
            items: {
                visible: item_visible,
                start: item_start
            },
            swipe: {
                onMouse: true,
                onTouch: true
            },
            scroll: {
                items: 1,
                duration: 700,
                pauseOnHover: 'resume'
            },
            prev: {
                button: '.header-two-prev',
                onBefore: function( data ) {
                    $('#header-two li').not(data.items.visible[item_animation]).find('.slider-border, .title, .slider-border, .border')
                        .animate({
                            opacity: 0
                        }, 100);
                    
                    $(data.items.visible[item_animation]).find('.slider-border, .title, .slider-border, .border')
                        .animate({
                            opacity: 1
                        }, 1500);
                },
            },
            next: {
                button: '.header-two-next',
                onBefore: function( data ) {
                    $('#header-two li').not(data.items.visible[item_animation]).find('.slider-border, .title, .slider-border, .border')
                        .animate({
                            opacity: 0
                        }, 100);
                    
                    $(data.items.visible[item_animation]).find('.slider-border, .title, .slider-border, .border')
                        .animate({
                            opacity: 1
                        }, 1500);
                },
            },
            pagination: {
                    container: '.header-two-pager',
                    deviation: 1
            }
        });

        $(".header-two-box-prev, .header-two-box-next").css({
            width: Math.round(($("body").width() - $("#header-two li:first-child").width()) / 2),
            height: $("#header-two li:first-child").height()
        });
        
        $(".header-two-prev").css("left", Math.round(($("body").width() - $("#header-two li:first-child").width()) / 2 - 70));
        $(".header-two-next").css("right", Math.round(($("body").width() - $("#header-two li:first-child").width()) / 2 - 70));
    }
    
    /* End Header Two Slider */
    
    /* Header Third Slider */
    $(".header-third-box-prev").click(function(){
        $("#header-third").trigger('prev')
    });
    $(".header-third-box-next").click(function(){
        $("#header-third").trigger('next')
    });

    if ($("#header-third")[0]) {
        if ($("#header-third li").length < 4) {
            $($("#header-third li")).each(function(){
                $(this).clone().appendTo("#header-third");
            });
        }
        $('#header-third').carouFredSel({
            width: '100%',
            height: 'auto',
            responsive: item_responsive,
            items: {
                    visible: item_visible,
                    start: item_start
            },
            swipe: {
                onMouse: true,
                onTouch: true
            },
            scroll: {
                    items: 1,
                    duration: 1000,
                    pauseOnHover: true
            },
            prev: {
                button: '.header-third-prev',
                onBefore: function( data ) {
                    $('#header-third li').not(data.items.visible[item_animation]).find('.slider-border, .title')
                        .animate({
                            opacity: 0
                        });
                    
                    $(data.items.visible[item_animation]).find('.slider-border, .title')
                        .animate({
                            opacity: 1
                        });
                }
            },
            next: {
                button: '.header-third-next',
                onBefore: function( data ) {
                    $('#header-third li').not(data.items.visible[item_animation]).find('.slider-border, .title')
                        .animate({
                            opacity: 0
                        });
                    
                    $(data.items.visible[item_animation]).find('.slider-border, .title')
                        .animate({
                            opacity: 1
                        });
                }
            },
            pagination: {
                    container: '.header-third-pager',
                    deviation: 1
            }
        });
        
        $('#header-third li ul').each(function () {
            $(this).carouFredSel({
                height: 'auto',
                items: 2,
                scrol: 2,
                swipe: {
                    onMouse: true,
                    onTouch: true
                },
                auto: false,
                pagination: $(this).closest("li").find(".nested-header-pager")
            });
        });
        
        var full_width = $("body").width();
        $(".header-third-box-prev, .header-third-box-next").css({
            width: Math.round(($("body").width() - $("#header-third li:first-child").width()) / 2),
            height: $("#header-third li:first-child").height()
        });
        
        $(".header-third-prev").css("left", Math.round(($("body").width() - $("#header-third li:first-child").width()) / 2 - 70));
        $(".header-third-next").css("right", Math.round(($("body").width() - $("#header-third li:first-child").width()) / 2 - 70));
        
        $("#header-third .nested-header-box")
            .bind("mousedown mousemove mouseup mouseleave touchstart touchmove touchend touchcancel", function (e) { e.stopPropagation() });
    }
    
    /* End Header Third Slider */

    $('input, select').styler();

        if ($("body").width() > 768) {
            $(".top-main-navigation > .dropdown, .nav-top-right .dropdown").hover(function(){
                $(this).addClass("open");
            }, function(){
                $(this).removeClass("open");
            });
            
            $(".top-main-navigation .dropdown .dropdown").hover(function(){
                $(".top-main-navigation .dropdown .dropdown").removeClass("open");
                $(this).addClass("open");
            }, function(){
                var this_item_dropdown = $(this);
                setTimeout(function(){
                    this_item_dropdown.removeClass("open");
                }, 200);
            });

            $(".main-menu .dropdown").hover(function(){
                $(this).addClass("open");
            }, function(){
                $(this).removeClass("open");
            });
        } else {
            $(".top-main-navigation .dropdown > a, .top-main-navigation .dropdown .dropdown > a").click(function(){
                if ($(this).parent().hasClass("open")) {
                    $(this).parent().removeClass("open");
                    return false;
                } else {
                    $(this).parent().addClass("open");
                    return false;
                }
            });
        }

/* Search Block Animation */

    /* Top Search Block */
    $("html").click(function(){
        if ($(".top-search-query").hasClass('focused')) {hide_top_search_input();}
    });
    $(".top-search-query").click(function(e){
        e.stopPropagation();
    });
    $(".top-search-button").click(function(){
        var top_search_input = $(".top-search-query");

        if (top_search_input.is(".focused")) {
            if (top_search_input.val().length > 0) {
                top_search_input.addClass("submitted");
                return true;
            }
            else
                return false;
        }
        else {
            return false;
        }
    });

    $(".top-search-button").click(function(){
        if (!$(".top-search-query").hasClass('focused')) {
            $(".top-search").addClass("open");
            $(".top-search").animate({width: $('body').width() < 440 ? "212px" : "362px"}, 200);
            $(".top-search-query").animate({width: $('body').width() < 440 ? "150px" : "300px", padding: "22px 14px"}, 200, function () {
                $(this).addClass('focused');
                $(this).stop(true);
                $(".top-search-query").focus();
            });
        } else {
            hide_top_search_input();
        }
    });

    function hide_top_search_input() {
        setTimeout(function () {
            if ($(".top-search-query").is(".submitted"))
                return;
            
            $(".top-search-query").removeClass("inline-block");
            $(".top-search").animate({width: $('body').width() < 440 ? "62px" : "62px"}, 200);
            $(".top-search-query").animate({width: "0", padding: "0"}, 200, function () {
                $(this).removeClass('focused');
                $(this).stop(true);
                $(".top-search").removeClass("open");
            });
        }, 150);
    }
    /* End Top Search Block */


    $(".search-button").click(function(){
        var search_input = $(".search-query");

        if (search_input.is(".focused")) {
            if (search_input.val().length > 0) {
                search_input.addClass("submitted");
                return true;
            }
            else
                return false;
        }
        else {
            search_input.focus();
            return false;
        }
    });

    function hide_search_input() {
        setTimeout(function () {
            if ($(".search-query").is(".submitted"))
                return;
            
            $(".search-query").removeClass("inline-block");
            if ($('body').width() < 768) {
                $(".search-query").animate({width: "0", right: 68, padding: 0}, 300, function () {
                    $(this).removeClass('focused');
                    $(".search-query").css("visibility", "hidden");
                    $(this).stop(true);
                });
            } else {
                $(".search-query").animate({width: "0", right: 52, padding: 0}, 300, function () {
                    $(this).removeClass('focused');
                    $(".search-query").css("visibility", "hidden");
                    $(this).stop(true);
                });
            }
        }, 150);
    }

    $(".search-button").click(function(){
        if (!$(".search-query").hasClass('focused')) {
            $(".search-query").addClass("inline-block");
            if ($('body').width() < 768) {
                $(".search-query").css("visibility", "visible");
                $(".search-query").animate({width: $(".main-menu .navbar").width() - 54, padding: "4px 14px", right: 68}, 300, function () {
                    $(this).addClass('focused');
                    $(this).stop(true);
                });    
            } else {
                $(".search-query").css("visibility", "visible");
                $(".search-query").animate({width: $('body').width() < 1024 ? "595px" : "916px", right: "52px", padding: "17px 14px"}, 300, function () {
                    $(this).addClass('focused');
                    $(this).stop(true);
                });
            }
        } else {
            hide_search_input();
        }
    });
    $("html").click(function(){
        if ($(".search-query").hasClass('focused')) {hide_search_input();}
    });
    $(".search-query").click(function(e){
        e.stopPropagation();
    });

    $(".caroufredsel_wrapper").mousedown(function () { $(".search-query").blur(); });
/* End Search Block Animation */    

/* Carousel Product */
    if ($('.carousel-grid .product-grid, .carousel-grid .content-list')[0]) {
        var scroll;
        $("body").width() < 567 ? scroll = 1 : scroll = true;
        $('.carousel-grid .product-grid, .carousel-grid .content-list').each(function () {
            $(this).carouFredSel({
                height: 'auto',
                auto: false,
                swipe: {
                    onMouse: true,
                    onTouch: true
                },
                scroll: scroll,
                pagination: $(this).closest(".carousel-grid").find(".carousel-pager"),
                prev: {
					button: $(this).closest(".carousel-grid").find(".carousel-prev"),
				},
                next: {
					button: $(this).closest(".carousel-grid").find(".carousel-next"),
				}
            });
        });
    }

/* End Carousel Product */

/* Product Zoom */

    if ($.fn.elevateZoom && $("body").width() > 768) {
        $("#main-product-image").elevateZoom({
            gallery:'product-images', 
            cursor: 'crosshair',
            galleryActiveClass: 'active',
            zoomWindowWidth: 343,
            zoomWindowHeight: 457,
            borderSize: 0,
        }); 
    }

/* End Product Zoom */

/* Product Images Carousel */
    
    if ($("#product-images")[0]) {
        $('#product-images').carouFredSel({
            height: 'auto',
            width: 276,
            items: 3,
            scroll: 1,
            auto: false,
            swipe: {
                onMouse: true,
                onTouch: true
            },
            prev: ".product-prev",
            next: ".product-next"
        });
    }

/* End  Product Images Carousel */

/* Related Product Sidebar */

    if ($(".relate-product-sidebar .list")[0]) {
        $('.relate-product-sidebar .list').carouFredSel({
            height: 'auto',
            scroll: 1,
            auto: false,
            swipe: {
                onMouse: true,
                onTouch: true
            },
            pagination: ".relate-product-sidebar-pager"
        });
    }

/* End Related Product Sidebar */

/* Single Project Pager*/
    $(".portfolio .item .carousel-grid").each(function () {
        var project_item = $(".carousel-pager", this).width() / $(".carousel-pager > a", this).length;
        $(".carousel-pager > a", this).css("width", project_item);
    });
/* End Single Project  Pager*/

/* Galery */

    $('.on-page-carousel a').modal('hide');

    var column_width;
    if ($("body").width() > 979) {
        column_width = 252;
    } else if ($("body").width() < 979 && $("body").width()  >= 768) {
        column_width = 195;
    }
    if ($.fn.masonry && $("body").width() > 767) {
        $('.gallery').imagesLoaded(function(){
            $('.gallery').masonry({
                itemSelector: '.image',
                columnWidth: column_width,
                isAnimated: true
            });
            $('.gallery.third').masonry({
                itemSelector: '.image',
                columnWidth: 336,
                isAnimated: true
            });
        });
    }


    $('.careers .image a, .gallery .image a').click(function() {
        var image_id = this.href.split('#').pop();
        $("#gallery-modal").css({display: 'block'});
        $('.gallery-carousel .content-list').trigger('slideTo', '#' + this.href.split('#').pop());
        $('.gallery-carousel-thumb a').removeClass('selected');
        $(".gallery-carousel-thumb a[href$='#" + this.href.split('#').pop() + "']").addClass('selected');
        setTimeout(function () {
            $("#gallery-modal").modal('show');
        }, 130);
        return false;
    });

    $('.gallery-carousel-thumb a').click(function() {
        $('.gallery-carousel .content-list').trigger('slideTo', '#' + this.href.split('#').pop());
        $('.gallery-carousel-thumb a').removeClass('selected');
        $(this).addClass('selected');
        return false;
    });

    $('.gallery-carousel .content-list').bind("slideTo.cfs next.cfs prev.cfs", function(e, a) {
        $(this).triggerHandler('currentPage', function (index) {
            $('.gallery-carousel-thumb a').removeClass('selected');
            $($('.gallery-carousel-thumb a')[index]).addClass('selected');
        });
    });

/* End Galery */

/* Transformation tabs in accordion */
if  ($("body").width() < 568) {
    $(".tab-content .tab-pane").each(function(){
        $(this).removeClass("fade");
        $(this).removeClass("in");
        $(this).appendTo($(".nav-tabs a[href$='" + this.id + "']").closest("li"));
        $(".nav-tabs > li").removeClass("active");
    });
}

/* End Transformation tabs in accordion */
    
};

$(window).resize(function(){
    var item_visible,
        item_start,
        item_animation,
        item_responsive;
    $("body").width() > 767 ? item_visible = 3 : item_visible = 1;
    $("body").width() > 767 ? item_start = 0 : item_start = -1;
    $("body").width() > 767 ? item_animation = 1 : item_animation = 0;
    $("body").width() > 767 ? item_responsive = false : item_responsive = true;

    theme_script();
});
$(document).ready(function(){
    var item_visible,
        item_start,
        item_animation,
        item_responsive;
    $("body").width() > 767 ? item_visible = 3 : item_visible = 1;
    $("body").width() > 767 ? item_start = 0 : item_start = -1;
    $("body").width() > 767 ? item_animation = 1 : item_animation = 0;
    $("body").width() > 767 ? item_responsive = false : item_responsive = true;

    theme_script();

/* Menu Up Dropdown */

$(".top-main-navigation > .dropdown > ul > .dropdown").each(function(){
    var items_inside = $("> ul > li", this);

    if ($(this).index() + items_inside.length > 17) {
        if (items_inside.length < 17) {
            $(this).addClass("show-top");
        }
    }
});

/* End Menu Up Dropdown */

/* Category Menu Dropdown */

$(".dropdown-section").click(function(){
    if (!$(this).hasClass("open")){
        $(this).addClass("open");
    } else {
        $(this).removeClass("open");
    }
    return false;
});

/* End Category Menu Dropdown */

    $(".minicolors").minicolors({
        defaultValue: "#fdf651",
        change: function(hex){
            $.cookie("color", "#222222");
            $.cookie("cart_button", "#222222");
            $.cookie("wishlist", "url('../img/black-wishlist-icon.png') no-repeat left 50%");
            $.cookie("compare", "url('../img/black-compare-icon.png') no-repeat left 50%");
            $.cookie("search_button", "rgba(0, 0, 0, .7)");
            $.cookie("border-color", "rgba(0, 0, 0, .3)");
            $.cookie("link-color", hex);
            changeColor(hex);
        }
    });

/* Portfolio Filter */

    var $filterType = $('#filterOptions a.active').attr('class');
    var $holder = $('.ourHolder');

    var $data = $holder.clone();

    $('#filterOptions a.category').click(function(e) {

        $('#filterOptions a.category').removeClass('active');

        var $filterType = $(this).attr('data-target');
        $(this).addClass('active');

        if ($filterType == 'all') {
            var $filteredData = $data.find('div.item-p');
        } 
        else {
            var $filteredData = $data.find('div.item-p[data-type=' + $filterType + ']');
        }

        $holder.quicksand($filteredData, {
            duration: 500,
            easing: 'easeInOutQuad'
        });
        return false;
    });

/* End Portfolio Filter */

/* Theme customizer */

    /* Customizer Panel */
    $("#options-panel .button").click(function(){
        if ($("#options-panel").hasClass("open")) {
            $("#options-panel").animate({left: -360}, 300, function(){
                $(this).removeClass("open");
            });
        } else {
            $("#options-panel").animate({left: 0}, 300, function(){
                $(this).addClass("open");
            });
        }
    });
    /* End Customizer Panel */

    /* Color Picker */
    var style = $('<style type="text/css" id="theme_color" />').appendTo('head');

    $(".color-scheme").click(function(){
        $.cookie("search_button", "rgba(0, 0, 0, .7)");
        $.cookie("border-color", "rgba(0, 0, 0, .3)");
        $.cookie("color", "#222222");
        $.cookie("cart_button", "#222222");
        $.cookie("wishlist", "url('../img/black-wishlist-icon.png') no-repeat left 50%");
        $.cookie("compare", "url('../img/black-compare-icon.png') no-repeat left 50%");
        $.cookie("button_color", $(this)[0]['attributes']['data-color']['value']);
        $.cookie("link-color", $(this)[0]['attributes']['data-color']['value']);
        changeColor($(this)[0]['attributes']['data-color']['value']);
    });

    $(".reset-theme-color").click(function(){
        $.cookie("color", "#ffffff");
        $.cookie("button_color", "#ad8020");
        $.cookie("search_button", "#504f3b");
        $.cookie("border-color", "#a29f43");
        $.cookie("theme_color_hex", "#fdf651");
        $.cookie("cart_button", "#77743a");
        $.cookie("wishlist", "url('../img/yellow-wishlist-icon.png') no-repeat left 50%");
        $.cookie("compare", "url('../img/yellow-compaire-icon.png') no-repeat left 50%");
        $.cookie("link-color", "#ad8020");
        $(".minicolors")[1]['value'] = $.cookie("theme_color_hex");
        $(".minicolors-swatch span").css("background-color", "#fdf651");
        changeColor($.cookie("theme_color_hex"));
        return false;
    });

    if ($.cookie("theme_color_hex")){
        $(".minicolors")[0]['value'] = $.cookie("theme_color_hex");
        changeColor($.cookie("theme_color_hex"));
    }

    function changeColor(hex) {
        $.cookie("theme_color_hex", hex, { expires: 7 });
        var search_button = $.cookie("search_button");
        var border_color = $.cookie("border-color");
        var color = $.cookie("color");
        var button_color = $.cookie("button_color");
        var cart_button = $.cookie("cart_button");
        var wishlist = $.cookie("wishlist");
        var compare = $.cookie("compare");
        var link_color = $.cookie("link-color");
        var rgba = parseInt(hex.substring(1), 16);
        var r = (rgba & 0xff0000) >> 16;
        var g = (rgba & 0x00ff00) >> 8;
        var b = rgba & 0x0000ff;
        style.html(
            '.main-menu .navbar, #options-panel .button, .product-grid .product-caption .inner, \n\
            .header-two-next:hover, .header-two-box-next:hover ~ .header-two-next, \n\
            .header-two-prev:hover, .header-two-box-prev:hover ~ .header-two-prev, .carousel-pager.one a:hover, \n\
            .carousel-next:hover, .carousel-prev:hover, .product-grid .product-caption .price, \n\
            .header-two-box .header-two-pager a.selected, .package .label, \n\
            .services .small .icon, .carousel-grid.brands.services li .icon, .portfolio .image .hover, \n\
            .container.newsletter-bottom-block .image, .baner-new-next:hover, .baner-new-prev:hover, #baner-new-pager > a.selected, \n\
            #sidebar-baners-pager > a.selected, .header-third-box .header-third-pager a.selected, \n\
            .header-third-prev:hover, .header-third-next:hover, \n\
            .header-third-box-next:hover ~ .header-third-next, .header-third-box-prev:hover ~ .header-third-prev, \n\
            #header-third .nested-header-box .slider-header .right a:hover, \n\
            .btn-ads, \n\
            .newsletter .info .item .label, .header-block.plan .image-label.mac, .price-package .label, \n\
            .toolbar .view-mode a.list:hover, .toolbar .view-mode a.grid:hover, .pagination ul li a:hover, \n\
            .about-us-page .about-us-page-slider .carousel-prev:hover, .about-us-page .about-us-page-slider .carousel-next:hover, \n\
            .progress .bar, .team-list .item .static-hover .description, .portfolio .portfolio-one .carousel-prev:hover, .portfolio .portfolio-one .carousel-next:hover, \n\
            .portfolio .item .carousel-pager a.selected, .portfolio .item .carousel-grid .carousel-prev:hover, .portfolio .item .carousel-grid .carousel-next:hover, \n\
            .portfolio .page-title .left:hover, .portfolio .page-title .right:hover, .portfolio .page-title .remove:hover, \n\
            .gallery-carousel .carousel-next:hover, .gallery-carousel .carousel-prev:hover, .services .large .icon, \n\
            #options-panel .layout a.active, #options-panel .layout a:hover, .mpc_ls .mpc_ls_ui .mpc_ls_timer, \n\
            .read_more, .mpc_ls_slides_thumbs.style01 .mpc_ls_thumb.active .mpc_ls_thumb_inside, \n\
            .mpc_ls_thumb:hover.active .mpc_ls_thumb_inside, .mpc_ls_image_thumb, .btn-ads span{background-color: ' + hex + '!important;}' + 

            '.navbar .nav .dropdown.open .dropdown-section a:hover, .county-days, .county-hours, .county-minutes, .county-seconds, \n\
            .navbar .nav li.dropdown.open .dropdown-menu.one a:hover{color: ' + hex + ';}' + 

            '.product-grid .product-caption .inner, .product-grid .product-caption .price, \n\
            .portfolio .image .hover, .team-list .item .static-hover .description, \n\
            .mpc_ls_slides_thumbs.style01 .mpc_ls_thumb.active .mpc_ls_thumb_inside, \n\
            .mpc_ls_slides_thumbs.style01 .mpc_ls_thumb:hover.active .mpc_ls_thumb_inside{background-color: rgba(' + r + ', ' + g + ', ' + b + ', 0.9)!important;}' + 

            '#header-third .nested-header-box .slider-header .right a.selected:hover, \n\
            .carousel-pager.one a.selected:hover{background-color: #444444!important;}' + 
            '.search-button{background-color: ' + search_button + '}' + 
            '.main-menu .navbar .nav, #header-two .title .slider-button, #header-one .title .slider-button, \n\
            #header-third .title .slider-button, .header-four-box .image-box .main-img .title .slider-button, \n\
            .header-seven-box .image-box .main-img .title .slider-button, .btn-ads span, \n\
            .main-menu .navbar .btn-navbar, .read_more span{border-color: ' + border_color + '!important;}' + 
            'a, a:hover, a:focus{color: ' + link_color + '}' + 

            '.navbar .nav li.dropdown.open a.btn-info, .btn-info{background-color: ' + button_color + '; color: ' + color + ';}' + 
            '.navbar .nav li.dropdown.open a.btn-info:hover, .btn-info:hover{border-color: ' + button_color + '}' + 
            '.product-grid .product-caption .cart-button .compaire-link, .product-grid .product-caption .cart-button .wishlist-link{color: ' + cart_button + '}' + 
            '.product-grid .product-caption .cart-button .wishlist-link{background: ' + wishlist + ';}' + 
            '.product-grid .product-caption .cart-button .compaire-link{background: ' + compare + ';}');
    }

    /* End Color Picker */

if ($.cookie("layout") == "wide") {
    changeLayout($.cookie("layout"));
} else if ($.cookie("layout") == "boxed") {
    changeLayout($.cookie("layout"));
} else if ($.cookie("layout") == "standart") {
    changeLayout($.cookie("layout"));
}

function changeLayout(layout){
    item_visible
    if (layout == "wide") {
		$("body").removeClass("boxed");
        $(".options a").removeClass("active");
        $(".options .wide").addClass("active");
        $(".navbar-static-top, .second-top-box, .header-one-box, .header-two-box, .header-third-box, .header-four-box, .header-five-box, .header-six-box, .header-seven-box, .header-eight-box, .header-nine-box, .site-container").removeClass("boxed");
        $(".site-container").addClass("wide");
        $('#header-two, #header-third').trigger("configuration", ["items.visible", item_visible, true]);
    } else if (layout == "standart") {
		$("body").removeClass("boxed");
        $(".options a").removeClass("active");
        $(".options .standart").addClass("active");
        $(".navbar-static-top, .second-top-box, .header-one-box, .header-two-box, .header-third-box, .header-four-box, .header-five-box, .header-six-box, .header-seven-box, .header-eight-box, .header-nine-box, .site-container").removeClass("boxed wide");
        $('#header-two, #header-third').trigger("configuration", ["items.visible", item_visible, true]);
    } else if (layout == "boxed") {
		$("body").addClass("boxed");
        $(".options a").removeClass("active");
        $(".options .boxed").addClass("active");
        $(".navbar-static-top, .second-top-box, .header-one-box, .header-two-box, .header-third-box, .header-four-box, .header-five-box, .header-six-box, .header-seven-box, .header-eight-box, .header-nine-box, .site-container").removeClass("boxed wide");
        $(".navbar-static-top, .second-top-box, .header-one-box, .header-two-box, .header-third-box, .header-four-box, .header-five-box, .header-six-box, .header-seven-box, .header-eight-box, .header-nine-box, .site-container").addClass("boxed");
        $('#header-two, #header-third').trigger("configuration", ["items.visible", item_visible, true]);
    }

}

$("a.wide").click(function(){
    $.cookie("layout", "wide");
    changeLayout("wide");
    $(".wide .header-four-box").css("background", "url('" + $.cookie("pattern") + "')");
    return false;
});

$("a.standart").click(function(){
    $.cookie("layout", "standart");
    changeLayout("standart");
    return false;
});

$("a.boxed").click(function(){
    $.cookie("layout", "boxed");
    changeLayout("boxed");
    return false;
});

if ($.cookie("pattern")) {
    $("body").css("background", "url('" + $.cookie("pattern") + "')");
    $(".wide .header-four-box").css("background", "url('" + $.cookie("pattern") + "')");
}

$(".pattern-scheme").click(function(){
    $.cookie("pattern", $(this)[0]["attributes"]["data-pattern"]["value"]);
    $("body").css("background", "url('" + $.cookie("pattern") + "')");
    $(".wide .header-four-box").css("background", "url('" + $.cookie("pattern") + "')");
});

$(".images-bg").click(function(){
    $.cookie("pattern", $(this)[0]["attributes"]["data-image"]["value"]);
    $("body").css("background", $.cookie("pattern"));
    //$(".wide .header-four-box").css("background", "url('img/bg.png')");
});

if ($.cookie("pattern")) {
    $("body").css("background", $.cookie("pattern"));
}

$(".reset-theme-pattern").click(function(){
    $.cookie("pattern", "img/bg.png");
    $("body").css("background", "url('" + $.cookie("pattern") + "')");
    $(".wide .header-four-box").css("background", "url('" + $.cookie("pattern") + "')");
    return false;
});

/* Singlr Project Synchronise Carousel */
    $(".portfolio .item .carousel-grid.on-page-carousel .content-list").bind("slideTo.cfs next.cfs prev.cfs", function() {
        this.synchronisation = true;
        if (!$(".portfolio .item .carousel-grid.modal-carousel .content-list")[0].synchronisation)
            $(".portfolio .item .carousel-grid.on-page-carousel .content-list").triggerHandler("synchronise", [".portfolio .item .carousel-grid.modal-carousel .content-list", true, true, 0]);
        this.synchronisation = false;
    });

    $(".portfolio .item .carousel-grid.modal-carousel .content-list").bind("slideTo.cfs next.cfs prev.cfs", function() {
        this.synchronisation = true;
        if (!$(".portfolio .item .carousel-grid.on-page-carousel .content-list")[0].synchronisation)
            $(".portfolio .item .carousel-grid.modal-carousel .content-list").trigger("synchronise", [".portfolio .item .carousel-grid.on-page-carousel .content-list", true, true, 0]);
        this.synchronisation = false;
    });
/* End Singlr Project Synchronise Carousel */

/* Single Project Tooltip */

    $('.portfolio .page-title a').tooltip();

/* End Single Project Tooltip */

/* Countdown */

    if ($.fn.county){
        $('#count-down').county({ endDateTime: new Date('2013/12/27 10:00:00'), reflection: false});
    }

/* End Countdown */

/* End Theme customizer */

/* Accordion Style */

    $(".accordion-body.in").parent().addClass("active"); 

    $(".accordion-heading").click(function(){
        var accordion_group = $(this).parent();
        if (accordion_group.hasClass("active")) {
            accordion_group.removeClass("visible");
            accordion_group.removeClass("active");
        } else {
            $(".accordion-body.in").parent().removeClass("active"); 
            $(".accordion-body.in").parent().removeClass("visible");
            accordion_group.addClass("active");
            setTimeout(function(){
                accordion_group.addClass("visible");
            }, 500);
        }
    });

/* End Accordion Style */

/* Product Tab */

    $('#product_tab a').click(function (e) {
      e.preventDefault();
      $(this).tab('show');
    });

/* End Product Tab */

$(".team-list .item").click(function(){
    if ($(this).hasClass("open")) {
        $("body").hover();
        $(this).removeClass("open");
    } else {
        $(this).hover().addClass("open");
    }
});

$(".btn-ads").hover(function(){
	if ($(this).hasClass("flip")) {
		$(this).removeClass("flip");
	} else {
		$(this).addClass("flip");
	}
});

});
