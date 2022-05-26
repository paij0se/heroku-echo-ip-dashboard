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

      return acc;
    }, "");

    document.getElementById("countries").innerHTML = countryList;
    document.getElementById("countries").style.fontFamily = "Cascadia Code";
    document.getElementById("countries").style.color = "white";
    document.getElementById("countries").style.fontSize = "20px";
  });
