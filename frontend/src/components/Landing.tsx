'use client';

import { userAPI } from '@/lib/api/users';
import { Box, Typography } from '@mui/material';
import Script from 'next/script';
import { useRef, useEffect } from 'react';

declare global {
  interface Window {
    google: any;
  }
}

export default function LandingPageComponent() {
  const googleButtonRef = useRef<HTMLDivElement>(null);

  const handleCredentialResponse = async (response: any) => {
    console.log('Google Auth Response:', response);

    try {
      // Send Google JWT to YOUR backend with correct object structure
      const result = await userAPI.createUser({
        googleToken: response.credential
      });
      console.log('Login successful:', result);

      if (result) {
        // Trigger auth change event
        window.dispatchEvent(new CustomEvent('auth-changed', {
          detail: { authenticated: true, user: result }
        }));
      } else {
        console.error('Login failed');
      }
    } catch (error) {
      console.error('Login error:', error);
    }
  };

  useEffect(() => {
    // Initialize Google button when SDK is loaded and component mounts
    const initializeGoogleButton = () => {
      if (window.google && window.google.accounts && googleButtonRef.current) {
        try {
          window.google.accounts.id.renderButton(googleButtonRef.current, {
            type: 'standard',
            size: 'large',
            text: 'signin_with',
            theme: 'outline',
            logo_alignment: 'left',
            width: 280,
            height: 50
          });
          console.log('Google button rendered successfully');
        } catch (error) {
          console.error('Error rendering Google button:', error);
        }
      }
    };

    // Check if SDK is already loaded
    if (window.google && window.google.accounts) {
      initializeGoogleButton();
    } else {
      // Wait for SDK to load
      const checkSDK = setInterval(() => {
        if (window.google && window.google.accounts) {
          clearInterval(checkSDK);
          initializeGoogleButton();
        }
      }, 100);
      
      // Clean up interval after 10 seconds
      setTimeout(() => clearInterval(checkSDK), 10000);
    }
  }, []);

  return (
    <>
      <Script
        src="https://accounts.google.com/gsi/client"
        strategy="afterInteractive"
        onLoad={() => {
          console.log('Google Client ID:', process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID);
          const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID?.trim();
          
          if (!clientId || clientId === 'undefined') {
            console.error('Google Client ID is not set or invalid');
            return;
          }
          
          if (window.google && window.google.accounts) {
            try {
              // Initialize Google Identity Services
              window.google.accounts.id.initialize({
                client_id: clientId,
                callback: handleCredentialResponse,
                auto_select: false,
                cancel_on_tap_outside: true,
                itp_support: true,
              });
              
              console.log('Google SDK initialized successfully');
            } catch (error) {
              console.error('Error initializing Google SDK:', error);
            }
          } else {
            console.error('Google SDK not available');
          }
        }}
        onError={(error) => {
          console.error('Failed to load Google SDK:', error);
        }}
      />
      <Box sx={{ 
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        textAlign: 'center',
        py: 8,
        gap: 3
      }}>
        <Typography variant="h1" component="h1" gutterBottom sx={{ fontWeight: 'bold' }}>
          TrackBot
        </Typography>
                
        {/* Google's official rendered button */}
        <div 
          ref={googleButtonRef}
          style={{ 
            minHeight: '50px',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center'
          }}
        />
      </Box>
    </>
  );
}
