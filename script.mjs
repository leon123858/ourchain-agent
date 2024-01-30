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
    let address = "";
    // get new address
    while (true) {
        try {
            // wait 10 second
            await new Promise(resolve => setTimeout(resolve, 10 * 1000));
            const tmp = await getNewAddress();
            console.log(`address: ${address}`);
            address = tmp;
            break;
        } catch (e) {
            console.log(e);
        }
    }
    while (true) {
        if(address === "") {
            throw new Error("address is empty");
        }
        try {
            await mineBlock(address);
            // wait 10 second
            // await new Promise(resolve => setTimeout(resolve, 10 * 1000));
        } catch (e) {
            console.log(e);
        }
    }
}

main();
