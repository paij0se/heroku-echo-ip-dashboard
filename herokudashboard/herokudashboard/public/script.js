const json = fetch(`https://${location.hostname}/ip`);
json
  .then((r) => {
    return r.json();
  })
  .then((_) => {});

const countries = fetch(`https://${location.hostname}/ip/all`);
countries
  .then((r) => {
    return r.text();
  })
  .then((countries) => {
    const countryCount = countries.split("\n").reduce((acc, country) => {
      if (country in acc) {
        acc[country]++;
      } else {
        acc[country] = 1;
      }
      return acc;
    }, {});
    const sortedCountries = Object.keys(countryCount).sort((a, b) => {
      return countryCount[b] - countryCount[a];
    });
    const countryList = sortedCountries.reduce((acc, country) => {
      acc += `${country}: ${countryCount[country]}\n`;
      let ConutryFlags = `https://countryflagsapi.com/png/${country.replace(
        /"/g,
        ""
      )}`;

      let arr = [ConutryFlags];
      for (let i = 0; i < arr.length; i++) {
        document.getElementById(
          "country-flag"
        ).innerHTML += `<img src="${arr[i]}">`;
        let br = document.createElement("br");
        document.getElementById("country-flag").appendChild(br);
      }
      const ctx = document.getElementById("myChart").getContext("2d");

      const myChart = new Chart(ctx, {
        type: "doughnut",
        data: {
          labels: sortedCountries,
          datasets: [
            {
              label: "Country",
              backgroundColor: [
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
              ],
              borderColor: [
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
              ],
              borderWidth: 1,
              data: Object.values(countryCount),
            },
          ],
        },
      });
      const ctxBar = document.getElementById("myChart2").getContext("2d");
      const barChart = new Chart(ctxBar, {
        type: "bar",
        data: {
          labels: sortedCountries,
          datasets: [
            {
              label: "Countries From Where Ip Address Was Visited",
              backgroundColor: [
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
              ],
              borderColor: [
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
              ],
              borderWidth: 1,
              data: Object.values(countryCount),
            },
          ],
        },
      });

      return acc;
    }, "");

    document.getElementById("countries").innerHTML = countryList;
    document.getElementById("countries").style.fontFamily = "Cascadia Code";
    document.getElementById("countries").style.color = "white";
    document.getElementById("countries").style.fontSize = "20px";
  });
