let interval = null;

function initialize() {
    showToast();
    sync();
    interval = setInterval(sync, 1000);
    document.addEventListener("keypress", onKeyPress);
}

function sync() {
    syncTimer();
    syncEvents();
    syncVotes();
    syncResults();
}

function syncTimer() {

    const start = new Date(room.resetTs);
    const now = new Date();
    const diff = now - start;
    const mins = Math.floor(diff / 60000);
    const secs = Math.floor((diff/1000)%60);
    const timer = padZero(""+mins)+":"+padZero(""+secs);
    const elem = document.getElementById("timer");
    elem.innerHTML = timer;

    function padZero(s) {
        if (s.length < 2) {
            return "0" + s;
        }
        return s;
    }
}

function syncEvents() {
    fetch("/rooms/" + room.name + "/events")
        .then(r => r.json())
        .then(d => {
            if (shouldReveal(d) || shouldReset(d)) {
                reload();
            }
        });

        function shouldReveal(d) {
            return d.revealed && !room.revealed;
        }

        function shouldReset(d) {
            return d.resetTs > room.resetTs
        }
}

function syncVotes() {
    fetch("/rooms/" + room.name + "/userlist")
        .then(r => r.text())
        .then(s => {
            const el = document.getElementById("userlist");
            el.innerHTML = s;
        });
}

function syncResults() {
    if (room.revealed) {
        fetch("/rooms/" + room.name + "/results")
            .then(r => r.text())
            .then(s => {
                const el = document.getElementById("results");
                el.innerHTML = s;
            });
    }
}

function vote(v) {
    fetch("/rooms/" + room.name + "/vote", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: v
    });
}

function makeToast(name, action) {
    M.toast({
        html: "<span class='amber-text'>"
            + name + "</span>&nbsp;" + action,
        displayLength: 15000,
     });
}

function showToast() {

    if (room.revealed) {
        makeToast(room.revealedBy, "revealed the votes");
    } else if (room.resetBy) {
        makeToast(room.resetBy, "started a new story")
    }
}

async function reveal() {
    await fetch("/rooms/" + room.name + "/reveal", {
        method: "POST"
    });
    reload();
}

async function reset() {
    await fetch("/rooms/" + room.name + "/reset", {
        method: "POST"
    });
    reload();
}

function reload() {
    clearInterval(interval);
    document.removeEventListener("keypress", onKeyPress);
    // https://developer.mozilla.org/en-US/docs/Web/API/Location/reload#parameters
    window.location.reload(true);
}

function onKeyPress(e) {
    const key = e.key
    const vote = shortcuts.get(key)
    if (vote) {
        const input = document.querySelector(`input[accesskey="${key}"]`);
        if (input) {
            input.click();
        }
    }
}

function copyToClipboard(element) {
    navigator.clipboard.writeText(element.innerText)
        .then(() => makeToast('Room link', 'copied to clipboard!'))
}

document.addEventListener('DOMContentLoaded', function() {
    M.Tooltip.init(document.querySelectorAll('.tooltipped'), {});
});
