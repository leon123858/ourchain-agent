// const domain = "127.0.0.1:8080"
const domain = "172.16.238.10:8080"

async function getNewAddress() {
    const result = await fetch(`http://${domain}/get/newaddress`, {
        method: 'GET'
    });
    const json = await result.json();
    return json.data;
}

async function mineBlock(address) {
	const path = address ? `http://${domain}/block/generate?address=${address}` : `http://${domain}/block/generate/`;
    return await fetch(path, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    });
}

async function main() {
    // get new address
    const address = await getNewAddress();
    while (true) {
        try {
            await mineBlock(address);
            // wait 10 second
            await new Promise(resolve => setTimeout(resolve, 10 * 1000));
        } catch (e) {
            console.log(e);
        }
    }
}

main();
