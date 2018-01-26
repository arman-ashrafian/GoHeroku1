$(() => {

    let blue = false
    $("#title").click(() => {
        let color = blue ? "black" : "blue"
        blue = !blue
        $("#title").css("color", color)
    })

})