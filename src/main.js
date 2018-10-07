!(function() {
  const init = () => {
    getData();
  };
  const click = e => {
    let target = e.target;
    if ([...target.classList].includes("color-item")) {
      let questionColor = document.querySelector(".questionColor");
      if (target.dataset.color === questionColor.dataset.color) {
        alert("ВЕРНО, МОЛОДЕЦ!");
        getData();
      } else {
        alert("попробуй снова");
        getData();
      }
    }
  };

  const getData = () => {
    let content = document.querySelector(".content");
    let questionColor = document.querySelector(".questionColor");

    content.addEventListener("click", click);

    fetch("./api/colors")
      .then(response => {
        return response.json();
      })
      .then(data => {
        drawQuestions(data, content, questionColor);
      });
  };
  const createElement = element => {
    let div = document.createElement("div");
    div.classList = "color color-item";
    div.dataset.color = element.name;
    div.style.cssText = `background-color:${element.name}`;
    return div;
  };

  drawQuestions = (data, content, questionColor) => {
    content.innerHTML = "";
    data.sort((a, b) => (Math.random(0.5) > 0.5 ? 1 : -1));
    data.forEach(element => {
      let div = createElement(element);
      content.appendChild(div);
    });
    let rundomQ = Math.floor(Math.random() * data.length - 1);
    rundomQ = rundomQ < 0 ? 0 : rundomQ;
    console.log(rundomQ);
    questionColor.innerHTML =
      data[rundomQ].transcription + " / " + data[rundomQ].name;
    questionColor.dataset.color = data[rundomQ].name;
  };

  window.onload = () => {
    init();
  };
})();
