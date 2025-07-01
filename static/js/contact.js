// Contact Form Handler
document.addEventListener('DOMContentLoaded', function() {
    const contactForm = document.getElementById('contactForm');
    
    if (contactForm) {
        // Ambil CSRF token dari meta tag
        const csrfToken = document.querySelector('meta[name="csrf-token"]')?.getAttribute('content');

        // Paksa trigger event input untuk semua field saat halaman dimuat (atasi autofill)
        ['name', 'email', 'phone', 'message'].forEach(function(id) {
            const el = document.getElementById(id);
            if (el && el.value) {
                el.dispatchEvent(new Event('input', { bubbles: true }));
            }
        });

        contactForm.addEventListener('submit', function(e) {
            e.preventDefault();
            const submitBtn = contactForm.querySelector('button[type="submit"]');
            const originalText = submitBtn.innerHTML;
            submitBtn.disabled = true;
            submitBtn.innerHTML = '<span class="spinner-border spinner-border-sm me-2"></span>Mengirim...';

            setTimeout(function() {
                const formData = new FormData();
                formData.append('name', document.getElementById('name').value);
                formData.append('email', document.getElementById('email').value);
                formData.append('phone', document.getElementById('phone').value);
                formData.append('message', document.getElementById('message').value);

                fetch('/contact', {
                    method: 'POST',
                    body: formData,
                    headers: csrfToken ? { 'X-CSRF-Token': csrfToken } : {}
                })
                .then(async response => {
                    let data = {};
                    try {
                        data = await response.json();
                    } catch (e) {
                        // Jika response kosong atau bukan JSON, anggap sukses jika status 200
                        if (response.ok) {
                            data = { success: true, message: "Pesan berhasil diterima!" };
                        } else {
                            data = { success: false, message: "Terjadi kesalahan pada server" };
                        }
                    }
                    return data;
                })
                .then(data => {
                    if (data.success) {
                        contactForm.reset();
                        // Tampilkan modal sukses
                        document.getElementById('successModalMsg').textContent = data.message || 'Pesan berhasil dikirim!';
                        const modal = new bootstrap.Modal(document.getElementById('successModal'));
                        modal.show();
                    } else {
                        // Tampilkan modal error
                        document.getElementById('errorModalMsg').textContent = data.message || 'Terjadi kesalahan saat mengirim pesan';
                        const modal = new bootstrap.Modal(document.getElementById('errorModal'));
                        modal.show();
                    }
                })
                .catch(error => {
                    console.error('Error:', error);
                    alert('Terjadi kesalahan saat mengirim pesan');
                })
                .finally(() => {
                    // Re-enable button
                    submitBtn.disabled = false;
                    submitBtn.innerHTML = originalText;
                });
            }, 100); // Delay 100ms
        });
    }
    
    // Handle modal close and page reload
    const successModal = document.getElementById('successModal');
    if (successModal) {
        successModal.addEventListener('hidden.bs.modal', function() {
            window.location.reload();
        });
    }

    // Batasi input phone hanya angka
    const phoneInput = document.getElementById('phone');
    if (phoneInput) {
        phoneInput.addEventListener('input', function(e) {
            this.value = this.value.replace(/[^0-9]/g, '');
        });
    }
});

// Utility function to show success message
function showSuccessMessage() {
    if (window.innerWidth > 768) {
        const modal = new bootstrap.Modal(document.getElementById('successModal'));
        modal.show();
    } else {
        alert('Pesan berhasil dikirim!');
        window.location.reload();
    }
} 