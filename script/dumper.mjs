const domain = '127.0.0.1:8080';

async function dump(address, args) {
    const dumpUrl = `http://${domain}/get/contractmessage`;
    const result = await fetch(dumpUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            address: address,
            arguments: [...args],
        }),
    });
    return await result.json();
}

async function main() {
    const tryCount = 25;
    const arr = [];
    for (let i = 0; i < tryCount; i++) {
        arr.push(
            dump(
                'ad1cfe47f515f88b9dd1710ad29f194177ee953a5df5cac261a1a97302dc8184',
                []
            )
        );
    }
    const result = await Promise.all(arr);
    console.log(result);
    // get success percentage for result is "success" or "fail"
    const success = result.filter((r) => r.result === 'success').length;
    console.log(`success: ${success}, fail: ${tryCount - success}`);
}

main();
