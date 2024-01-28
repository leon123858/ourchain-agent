const domain = "172.16.238.10:8080"

async function mineBlock() {
	return await fetch(`http://${domain}/block/generate`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
	});
}

async function main() {
	 while (true) {
		try {
			await mineBlock();
			// wait 2 second
			await new Promise(resolve => setTimeout(resolve, 2*1000));
		} catch (e) {
			console.log(e);
		}
	}
}

main();
