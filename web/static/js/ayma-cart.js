(() => {
    const cartKey = "ayma_cart";

    const readCart = () => {
        try {
            const raw = localStorage.getItem(cartKey);
            return raw ? JSON.parse(raw) : [];
        } catch {
            return [];
        }
    };

    const writeCart = (items) => {
        localStorage.setItem(cartKey, JSON.stringify(items));
        updateHeaderCount(items);
    };

    const updateHeaderCount = (items = readCart()) => {
        const count = items.reduce((sum, item) => sum + item.qty, 0);
        document.querySelectorAll("[data-cart-count]").forEach((el) => {
            el.textContent = count;
        });
    };

    const addToCart = (product, qty) => {
        const items = readCart();
        const existing = items.find((item) => item.id === product.id);
        if (existing) {
            existing.qty += qty;
        } else {
            items.push({ ...product, qty });
        }
        writeCart(items);
    };

    const getQtyFromCard = (card) => {
        const input = card.querySelector(".qty");
        const value = input ? parseInt(input.value, 10) : 1;
        return Number.isFinite(value) && value > 0 ? value : 1;
    };

    const bindQuantityControls = () => {
        document.querySelectorAll(".quantity").forEach((wrapper) => {
            const input = wrapper.querySelector(".qty");
            const minus = wrapper.querySelector(".minus");
            const plus = wrapper.querySelector(".plus");
            if (!input || !minus || !plus) {
                return;
            }

            const syncState = () => {
                const value = parseInt(input.value, 10) || 1;
                input.value = value < 1 ? 1 : value;
                minus.disabled = input.value <= 1;
            };

            minus.addEventListener("click", () => {
                input.value = Math.max(1, (parseInt(input.value, 10) || 1) - 1);
                syncState();
            });

            plus.addEventListener("click", () => {
                input.value = (parseInt(input.value, 10) || 1) + 1;
                syncState();
            });

            input.addEventListener("change", syncState);
            syncState();
        });
    };

    const bindAddButtons = () => {
        document.querySelectorAll(".add_to_cart_button").forEach((button) => {
            button.addEventListener("click", (event) => {
                event.preventDefault();
                const card = button.closest(".product__wrapper");
                if (!card) {
                    return;
                }
                const qty = getQtyFromCard(card);
                const product = {
                    id: String(button.dataset.productId || ""),
                    name: button.dataset.productName || "Товар",
                    price: parseFloat(button.dataset.productPrice || "0"),
                    unit: button.dataset.productUnit || "",
                    image: "/static/ligopak/main_files/el000003041_1-300x300.jpg",
                };
                addToCart(product, qty);
            });
        });
    };

    const renderCart = () => {
        const container = document.querySelector("[data-cart-container]");
        const itemsBody = document.querySelector("[data-cart-items]");
        const totalEl = document.querySelector("[data-cart-total]");
        const emptyEl = document.querySelector("[data-cart-empty]");
        if (!container || !itemsBody || !totalEl || !emptyEl) {
            return;
        }

        const items = readCart();
        itemsBody.innerHTML = "";

        if (!items.length) {
            container.style.display = "none";
            emptyEl.style.display = "block";
            totalEl.textContent = "0 Бел.руб";
            return;
        }

        container.style.display = "";
        emptyEl.style.display = "none";

        let total = 0;
        items.forEach((item) => {
            const subtotal = item.price * item.qty;
            total += subtotal;
            const row = document.createElement("tr");
            row.className = "cart_item";
            row.dataset.cartId = item.id;
            row.innerHTML = `
                <td class="product-thumbnail">
                    <img src="${item.image}" alt="${item.name}">
                </td>
                <td class="product-name">
                    <div>${item.name}</div>
                    <div class="cart-item-price">${item.price.toFixed(2)} Бел.руб / ${item.unit}</div>
                </td>
                <td class="product-quantity">
                    <input type="number" class="cart-qty" min="1" value="${item.qty}">
                </td>
                <td class="product-subtotal">${subtotal.toFixed(2)} Бел.руб</td>
                <td class="product-remove"><a href="#" class="cart-remove">x</a></td>
            `;
            itemsBody.appendChild(row);
        });

        totalEl.textContent = `${total.toFixed(2)} Бел.руб`;

        itemsBody.querySelectorAll(".cart-remove").forEach((btn) => {
            btn.addEventListener("click", (event) => {
                event.preventDefault();
                const row = btn.closest("tr");
                if (!row) {
                    return;
                }
                const id = row.dataset.cartId;
                const next = readCart().filter((item) => item.id !== id);
                writeCart(next);
                renderCart();
            });
        });

        itemsBody.querySelectorAll(".cart-qty").forEach((input) => {
            input.addEventListener("change", () => {
                const row = input.closest("tr");
                if (!row) {
                    return;
                }
                const id = row.dataset.cartId;
                const next = readCart().map((item) => {
                    if (item.id === id) {
                        const qty = parseInt(input.value, 10);
                        return { ...item, qty: Number.isFinite(qty) && qty > 0 ? qty : 1 };
                    }
                    return item;
                });
                writeCart(next);
                renderCart();
            });
        });
    };

    document.addEventListener("DOMContentLoaded", () => {
        updateHeaderCount();
        bindQuantityControls();
        bindAddButtons();
        renderCart();
    });
})();

