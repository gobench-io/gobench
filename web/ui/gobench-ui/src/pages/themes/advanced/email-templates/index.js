import React from 'react'
import { Helmet } from 'react-helmet'

const AdvancedEmailTemplates = () => {
  return (
    <div>
      <Helmet title="Advanced / Email Templates" />
      <div className="row">
        <div className="col-lg-12">
          <h5 className="mb-4">
            <strong>Simple Email</strong>
          </h5>
          <div className="mb-5">
            <div
              style={{
                background: '#eceff4',
                padding: '50px 20px',
                color: '#514d6a',
                borderRadius: '5px',
              }}
            >
              <div style={{ maxWidth: '700px', margin: '0px auto', fontSize: '14px' }}>
                <table
                  cellPadding="0"
                  cellSpacing="0"
                  style={{ width: '100%', marginBottom: '20px', border: '0px' }}
                >
                  <tbody>
                    <tr>
                      <td style={{ verticalAlign: 'top' }}>
                        <img
                          src="resources/images/tf-logo.png"
                          style={{ height: '40px' }}
                        />
                      </td>
                      <td style={{ textAlign: 'right', verticalAlign: 'middle' }}>
                        <span style={{ color: '#a09bb9' }}>Some Description</span>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div style={{ padding: '40px 40px 20px 40px', background: '#fff' }}>
                  <table cellPadding="0" cellSpacing="0" style={{ width: '100%', border: '0px' }}>
                    <tbody>
                      <tr>
                        <td>
                          <p>Hi there,</p>
                          <p>
                            Sometimes you just want to send a simple HTML email with a simple design
                            and clear call to action.
                          </p>
                          <a
                            style={{
                              display: 'inline-block',
                              padding: '11px 30px 6px',
                              margin: '20px 0px 30px',
                              fontSize: '15px',
                              color: '#fff',
                              background: '#01a8fe',
                              borderRadius: '5px',
                            }}
                          >
                            Call To Action
                          </a>
                          <p>
                            This is a really simple email template. It&apos;s sole purpose is to get
                            the recipient to click the button with no distractions.
                          </p>
                          <p>Good luck! Hope it works.</p>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div
                  style={{
                    textAlign: 'center',
                    fontSize: '12px',
                    color: '#a09bb9',
                    marginTop: '20px',
                  }}
                >
                  <p>
                    Mediatec Software Inc., Abbey Road, San Francisco CA 94102
                    <br />
                    Don&apos;t like these emails?{' '}
                    <a
                      style={{
                        color: '#a09bb9',
                        textDecoration: 'underline',
                      }}
                    >
                      Unsubscribe
                    </a>
                    <br />
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-lg-12">
          <h5 className="mb-4">
            <strong>Email w/ Header</strong>
          </h5>
          <div className="mb-5">
            <div
              width="100%"
              style={{
                background: '#eceff4',
                padding: '50px 20px',
                color: '#514d6a',
                borderRadius: '5px',
              }}
            >
              <div style={{ maxWidth: '700px', margin: '0px auto', fontSize: '14px' }}>
                <table
                  cellPadding="0"
                  cellSpacing="0"
                  style={{ width: '100%', marginBottom: '20px', border: '0px' }}
                >
                  <tbody>
                    <tr>
                      <td style={{ verticalAlign: 'top' }}>
                        <img
                          src="resources/images/tf-logo.png"
                          style={{ height: '40px' }}
                        />
                      </td>
                      <td style={{ textAlign: 'right', verticalAlign: 'middle' }}>
                        <span style={{ color: '#a09bb9' }}>Some Description</span>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div style={{ padding: '40px 40px 20px 40px', background: '#fff' }}>
                  <table cellPadding="0" cellSpacing="0" style={{ width: '100%', border: '0px' }}>
                    <tbody>
                      <tr>
                        <td>
                          <h5
                            style={{
                              marginBottom: '20px',
                              color: '#24222f',
                              fontWeight: '600',
                            }}
                          >
                            Password Reset
                          </h5>
                          <p>
                            Seems like you forgot your password for  KIT Pro. If this is
                            true, click below to reset your password.
                          </p>
                          <div style={{ textAlign: 'center' }}>
                            <a
                              style={{
                                display: 'inline-block',
                                padding: '11px 30px 6px',
                                margin: '20px 0px 30px',
                                fontSize: '15px',
                                color: '#fff',
                                background: '#01a8fe',
                                borderRadius: '5px',
                              }}
                            >
                              Reset Password
                            </a>
                          </div>
                          <p>
                            If you did not forgot your password you can safely ignore his email.
                          </p>
                          <p>
                            Regards,
                            <br />
                            Mediatec Software
                          </p>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div
                  style={{
                    textAlign: 'center',
                    fontSize: '12px',
                    color: '#a09bb9',
                    marginTop: '20px',
                  }}
                >
                  <p>
                    Mediatec Software Inc., Abbey Road, San Francisco CA 94102
                    <br />
                    Don&apos;t like these emails?{' '}
                    <a style={{ color: '#a09bb9', textDecoration: 'underline' }}>Unsubscribe</a>
                    <br />
                    Powered by gobench KIT Pro
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col-lg-12">
          <h5 className="mb-4">
            <strong>Ecommerce Email</strong>
          </h5>
          <div className="mb-5">
            <div
              style={{
                width: '100%',
                background: '#eceff4',
                padding: '50px 20px',
                color: '#514d6a',
                borderRadius: '5px',
              }}
            >
              <div style={{ maxWidth: '700px', margin: '0px auto', fontSize: '14px' }}>
                <table
                  cellPadding="0"
                  cellSpacing="0"
                  style={{ width: '100%', marginBottom: '20px', border: '0px' }}
                >
                  <tbody>
                    <tr>
                      <td style={{ verticalAlign: 'top' }}>
                        <img
                          src="resources/images/tf-logo.png"
                          alt="gobench KIT Pro"
                          style={{ height: '40px' }}
                        />
                      </td>
                      <td style={{ textAlign: 'right', verticalAlign: 'middle' }}>
                        <span style={{ color: '#a09bb9' }}>Some Description</span>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div style={{ padding: '40px 40px 20px 40px', background: '#fff' }}>
                  <table cellPadding="0" cellSpacing="0" style={{ width: '100%', border: '0px' }}>
                    <tbody>
                      <tr>
                        <td>
                          <h2
                            style={{
                              marginBottom: '20px',
                              color: '#24222f',
                              fontWeight: '600',
                            }}
                          >
                            Thanks for Purchase!
                          </h2>
                          <p>
                            <span style={{ color: '#a09bb9' }}>Monday, Dec 28 2015 at 4:13 PM</span>
                          </p>
                          <br />
                          <h5
                            style={{
                              marginBottom: '20px',
                              color: '#24222f',
                              fontWeight: '600',
                            }}
                          >
                            Your Order #00002345
                          </h5>
                          <table
                            cellPadding="0"
                            cellSpacing="0"
                            style={{ width: '100%', border: '0px' }}
                          >
                            <tbody>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                >
                                  Apple iPhone 6S
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '10px 10px',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                >
                                  1
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '10px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                >
                                  $ 699.00
                                </td>
                              </tr>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  Data cable
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '10px 10px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  1
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '10px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  $ 9.98
                                </td>
                              </tr>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  Nueng Silver Case
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '10px 10px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  2
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '10px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  $ 17.49
                                </td>
                              </tr>
                              <tr style={{ color: '#a09bb9' }}>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  Subtotal
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '10px 10px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  4
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '10px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  $ 735.96
                                </td>
                              </tr>
                              <tr style={{ color: '#a09bb9' }}>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '0px 10px 10px 0px',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  Tax
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '0px 10px',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  10%
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '0px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  $ 73.60
                                </td>
                              </tr>
                              <tr style={{ color: '#a09bb9' }}>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '0px 10px 10px 0px',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  Shipping
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '0px 10px',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  &nbsp;
                                </td>
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '0px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '0px solid #d9d7e0',
                                  }}
                                >
                                  $ 9.99
                                </td>
                              </tr>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                >
                                  <span style={{ fontSize: '18px', fontWeight: 'bold' }}>
                                    Total
                                  </span>
                                </td>
                                <td
                                  style={{
                                    width: '10%',
                                    textAlign: 'center',
                                    padding: '10px 10px',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                />
                                <td
                                  style={{
                                    width: '20%',
                                    textAlign: 'right',
                                    padding: '10px 0px 10px 10px',
                                    whiteSpace: 'nowrap',
                                    borderTop: '3px solid #514d6a',
                                  }}
                                >
                                  <span style={{ fontSize: '18px', fontWeight: 'bold' }}>
                                    $ 876.96
                                  </span>
                                </td>
                              </tr>
                            </tbody>
                          </table>
                          <br />
                          <br />
                          <h5
                            style={{
                              marginBottom: '20px',
                              color: '#24222f',
                              fontWeight: '600',
                            }}
                          >
                            Your Details
                          </h5>
                          <table
                            cellPadding="0"
                            cellSpacing="0"
                            style={{ width: '100%', border: '0px' }}
                          >
                            <tbody>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '1px solid #d9d7e0',
                                    whiteSpace: 'nowrap',
                                    verticalAlign: 'top',
                                  }}
                                >
                                  Shipping To
                                </td>
                                <td
                                  style={{
                                    width: '50%',
                                    padding: '10px 0px 10px 10px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  Tony Stark
                                  <br />
                                  22 23rd Street
                                  <br />
                                  San Francisco
                                  <br />
                                  CA 94107
                                </td>
                              </tr>
                              <tr>
                                <td
                                  style={{
                                    textAlign: 'left',
                                    padding: '10px 10px 10px 0px',
                                    borderTop: '1px solid #d9d7e0',
                                    whiteSpace: 'nowrap',
                                    verticalAlign: 'top',
                                  }}
                                >
                                  Billed To
                                </td>
                                <td
                                  style={{
                                    width: '50%',
                                    padding: '10px 0px 10px 10px',
                                    borderTop: '1px solid #d9d7e0',
                                  }}
                                >
                                  Visa
                                  <br />
                                  Ending in *7643
                                  <br />
                                  Expiring 08/2020
                                </td>
                              </tr>
                            </tbody>
                          </table>
                          <br />
                          <br />
                          <p style={{ textAlign: 'center' }}>
                            Notice something wrong?{' '}
                            <a style={{ color: '#01a8fe', textDecoration: 'underline' }}>
                              Contact our support team
                            </a>{' '}
                            and we&apos;ll e happy to help.
                          </p>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div
                  style={{
                    textAlign: 'center',
                    fontSize: '12px',
                    color: '#a09bb9',
                    marginTop: '20px',
                  }}
                >
                  <p>
                    Mediatec Software Inc., Abbey Road, San Francisco CA 94102
                    <br />
                    Don&apos;t like these emails?{' '}
                    <a style={{ color: '#a09bb9', textDecoration: 'underline' }}>Unsubscribe</a>
                    <br />
                    Powered by gobench KIT Pro
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-lg-12">
          <h5 className="mb-4">
            <strong>Email w/ Action</strong>
          </h5>
          <div className="mb-5">
            <div
              width="100%"
              style={{
                background: '#eceff4',
                padding: '50px 20px',
                color: '#514d6a',
                borderRadius: '5px',
              }}
            >
              <div style={{ maxWidth: '700px', margin: '0px auto', fontSize: '14px' }}>
                <table
                  cellPadding="0"
                  cellSpacing="0"
                  style={{ width: '100%', marginBottom: '20px', border: '0px' }}
                >
                  <tbody>
                    <tr>
                      <td style={{ verticalAlign: 'top' }}>
                        <img
                          src="resources/images/tf-logo.png"
                          alt="gobench KIT Pro"
                          style={{ height: '40px' }}
                        />
                      </td>
                      <td style={{ textAlign: 'right', verticalAlign: 'middle' }}>
                        <span style={{ color: '#a09bb9' }}>Some Description</span>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div style={{ padding: '40px 40px 20px 40px', background: '#fff' }}>
                  <table cellPadding="0" cellSpacing="0" style={{ width: '100%', border: '0px' }}>
                    <tbody>
                      <tr>
                        <td>
                          <div
                            style={{
                              padding: '15px 30px',
                              background: '#46be8a',
                              borderRadius: '5px',
                              marginBottom: '20px',
                              color: '#fff',
                            }}
                          >
                            Success! Something good happened.
                          </div>
                          <div
                            style={{
                              padding: '15px 30px',
                              background: '#fb434a',
                              borderRadius: '5px',
                              marginBottom: '20px',
                              color: '#fff',
                            }}
                          >
                            Error! Something bad happened.
                          </div>
                          <div
                            style={{
                              padding: '15px 30px',
                              background: '#fff',
                              border: '1px solid #acb7bf',
                              borderRadius: '5px',
                              marginBottom: '20px',
                            }}
                          >
                            Information! Something neutral happened.
                          </div>
                          <p>Hi George,</p>
                          <p>Congratulations! Something good has appened.</p>
                          <div style={{ textAlign: 'center' }}>
                            <a
                              style={{
                                display: 'inline-block',
                                padding: '11px 30px 6px',
                                margin: '20px 0px 30px',
                                fontSize: '15px',
                                color: '#fff',
                                background: '#01a8fe',
                                borderRadius: '5px',
                              }}
                            >
                              Take Action Now
                            </a>
                          </div>
                          <p>Thanks for being great customer. Let it be!</p>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div
                  style={{
                    textAlign: 'center',
                    fontSize: '12px',
                    color: '#a09bb9',
                    marginTop: '20px',
                  }}
                >
                  <p>
                    Mediatec Software Inc., Abbey Road, San Francisco CA 94102
                    <br />
                    Don&apos;t like these emails?{' '}
                    <a style={{ color: '#a09bb9', textDecoration: 'underline' }}>Unsubscribe</a>
                    <br />
                    Powered by gobench KIT Pro
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AdvancedEmailTemplates
