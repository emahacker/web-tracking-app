<script>
    let socialTime = 0;
    let webTime = 0;

    let socialDomains = ['facebook.com', 'instagram.com', 'twitter.com']; // Aggiungi qui i domini dei social

    function startTracking() {
        setInterval(function() {
            let currentDomain = document.location.hostname;
            if (socialDomains.includes(currentDomain)) {
                socialTime++;
                console.log("Tempo sui Social: " + socialTime + " minuti");
            } else {
                webTime++;
                console.log("Tempo su Web: " + webTime + " minuti");
            }
        }, 60000); // Ogni 1 minuto
    }
</script>
