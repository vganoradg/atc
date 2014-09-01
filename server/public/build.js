var autoscroll = false;

var eventHandlers = {
  "0.0": function(msg) {
    writeLogs(msg.data, $("#build-log"));
  },
  "1.0": function(msg) {
    var eventMsg = JSON.parse(msg.data);

    switch(eventMsg.type) {
    case "log":
      processLogs(eventMsg.event);
      break;
    case "error":
      var errorSpan = $("<span>");
      errorSpan.addClass("error");
      errorSpan.text(eventMsg.event.message);

      $("#build-log").append(errorSpan);
    case "status":
      var currentStatus = $("#build-title").attr("class");

      // only transition from transient states; state may already be set
      // if the page loaded after build was done
      if(currentStatus != "pending" && currentStatus != "started") {
        break;
      }

      var status = eventMsg.event.status;

      $("#build-title").attr("class", status);
      $("#builds .current").attr("class", status + " current");

      if(status != "started") {
        $(".abort-build").remove();
      }

      break;
    }
  }
}

function processLogs(event) {
  var log;

  if(event.origin.type == "run") {
      log = $("#build-log");
  } else if(event.origin.type == "input") {
      log = $(".build-inputs .build-source[data-resource='"+event.origin.name+"'] .log");
  } else if(event.origin.type == "output") {
      log = $(".build-outputs .build-source[data-resource='"+event.origin.name+"'] .log");
  }

  if(!log) {
    return;
  }

  writeLogs(event.payload, log);
}

function writeLogs(payload, destination) {
  var sequence = ansiparse(payload);

  var ele;
  for(var i = 0; i < sequence.length; i++) {
    ele = $("<span>");
    ele.text(sequence[i].text);

    if(sequence[i].foreground) {
      ele.addClass("ansi-"+sequence[i].foreground+"-fg");
    }

    if(sequence[i].background) {
      ele.addClass("ansi-"+sequence[i].background+"-bg");
    }

    if(sequence[i].bold) {
      ele.addClass("ansi-bold");
    }

    destination.append(ele);
  }

  if (autoscroll) {
    $(document).scrollTop($(document).height());
  }
}

function streamLog(uri) {
  var ws = new WebSocket(uri);

  var eventHandler;
  ws.onmessage = function(event) {
    if(!eventHandler) {
      var versionMsg = JSON.parse(event.data);

      if(versionMsg.version) {
        eventHandler = eventHandlers[versionMsg.version];
      }
    } else {
      eventHandler(event);
    }
  };
}

function scrollToCurrentBuild() {
  var currentBuild = $("#builds .current");
  var buildWidth = currentBuild.width();
  var left = currentBuild.offset().left;

  if((left + buildWidth) > window.innerWidth) {
    $("#builds").scrollLeft(left - buildWidth);
  }
}

$(document).ready(function() {
  var title = $("#build-title");

  if (title.hasClass("pending") || title.hasClass("started"))
    autoscroll = true;

  $(window).scroll(function() {
    var scrollEnd = $(window).scrollTop() + $(window).height();

    if (scrollEnd >= ($(document).height() - 16)) {
      autoscroll = true;
    } else {
      autoscroll = false;
    }
  });

  $("#builds").bind('mousewheel', function(e){
    if (e.originalEvent.deltaX != 0) {
      $(this).scrollLeft($(this).scrollLeft() + e.originalEvent.deltaX);
    } else {
      $(this).scrollLeft($(this).scrollLeft() - e.originalEvent.deltaY);
    }

    return false;
  });

  if ($(".resource-body").size() > 1)
    $(".resource-body").hide();

  $(".resource-header").click(function() {
    $(this).parent().find(".resource-body").toggle();
  });

  scrollToCurrentBuild();
});
