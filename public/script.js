const json = fetch(`https://${location.hostname}/ip`);
json
  .then((r) => {
    return r.json();
  })
  .then((ip) => {
    console.log(ip);
    document.getElementById("ip").innerHTML = ip.query;
    document.getElementById("city").innerHTML = ip.city;
    document.getElementById("country").innerHTML = ip.country;
    document.title = ip.country;
    document.getElementById(
      "flag"
    ).src = `https://countryflagsapi.com/png/${ip.country}`;
  });
